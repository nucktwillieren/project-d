# project-d

## Goals of this Project
- Dcard 每天午夜都有大量使用者湧入抽卡，為了不讓伺服器過載，請設計一個 middleware：
- 限制每小時來自同一個 IP 的請求數量不得超過 1000
- 在 response headers 中加入剩餘的請求數量 (X-RateLimit-Remaining) 以及 rate limit 歸零的時間 (X-RateLimit-Reset)
- 如果超過限制的話就回傳 429 (Too Many Requests)
- 可以使用各種資料庫達成
- Deadline: 4/5 23:59

## Analysis
### 限制每小時來自同一個 IP 的請求數量不得超過 1000
- 如何才會讓這個請求數量1000判斷比較合理?
  - 是否應當將載入static files or media files的請求也算入? 這樣是否合理? middleware介入的位置? 或者何種架構可以避免算入這種情況? 使用特定且和此middleware隔離的CDN Server?
  - 除了IP，是否應有更多判斷考量? 有些用戶可能是使用租屋處的網路，若是在租屋處共有的Public IP的情況下使用，假使某段時間該租屋處的IP使用正好加起來超過1000，並封鎖，此情況是否合理? 
  - 所謂每個小時?
    - 移動的一個小時?
    - 每個小時刷新一次總數?

### 在 response headers 中加入剩餘的請求數量 (X-RateLimit-Remaining) 以及 rate limit 歸零的時間(X-RateLimit-Reset)
- remaining = max - count 
- time_remaining = TTL

### 如果超過限制的話就回傳 429 (Too Many Requests)
- Abort in middleware

### 可以使用各種資料庫達成
- in-memory database或許是一個比較好的選擇

## Basic Dependencies
- docker-compose 1.28.0+
- docker 20.10.5+

## Setup
先Git Clone下來 
```
git clone https://github.com/nucktwillieren/project-d.git
```

進到project-d資料夾
```
cd project-d
```

使用docker-compose執行這個專案的所有服務(要注意docker-compose版本要1.28.0以上，因為使用了service profile的功能)
```
docker-compose --profile all up --build 
```

Build完之後，因為Nginx的關係，他能將port 80導到frontend的port上，因此直接在瀏覽器輸入 http://127.0.0.1 或 http://localhost 就可以連上frontend

## 架構說明
### Docker Containers
- project-d-nginx
  - Reverse Proxy and Load Balancer
  - 在此專案中沒有使用Load Balancer
- project-d-qcard-frontend
  - Built by using React.js
  - 在此專案中連接到project-d-qcard-go中取得後端資訊
  - 在此專案中是直接使用npm start開啟的dev server
  - 在一般專案中可以將npm build後產生的static直接使用nginx將static files發出
  - 在此專案中，這個frontend的功能如下
    - 創建使用者(註冊，登入前點擊右上的Sign Up)
    - 使用者驗證(登入。在登入前點擊右上的Sign In)
    - 使用者登出(登入後，點擊右上的Sign Out)
    - 查看個人檔案(登入後，點擊自己的名字)
    - 抽卡(登入後，點擊Top Nav中的Pick A Card)
      - ※備註：整個系統中需要至少兩個使用者以上才可以抽卡，因此若要測試，就需要註冊兩個
    - 創建貼文類型(創建看版)
    - 進入貼文類型(進入看版)
- project-d-qcard-go
  - 本專案連接postgreSQL的api server
  - 本專案使用gin-gonic
  - 透過grpc連接project-d-xlimit-grpc
- project-d-qcard-pg
  - 儲存所有使用者資料
- project-d-xlimit-redis
  - 儲存xlimit-rate的count，即題目所要求
- project-d-xlimit-grpc
  - 連接project-d-xlimit-redis，並透過grpc server跟外處溝通
  - 在此container的資料夾中，
    -  cmd資料夾的client可以做單一identity的request test
    -  cmd資料夾的random-create可以隨機產生identiy進行request test
    -  cmd資料夾的get可以取得所有key跟value   
- project-d-migration_tool
  - 利用python scripts快速的建立Table與column及constraint

## 解釋
### 為何先透過grpc的一層封裝將redis DB藏在grpc server後方
- 在分散式的系統中，我們可能會需要把關存取redis的關口，而grpc server恰好可以做到這件事情
- grpc速度比restful快
### 為何在redis中是直接使用SET而不是HSET
- 因為SET有EX(expire time)的選項
- key儲存為該個人資訊，利用EX作為倒數計時
