import "./LeftNav.css"
import axios from "axios"
import { Button, Modal, Form } from 'react-bootstrap';
import { useState, useEffect } from "react";
import { useSelector } from "react-redux"

export function CreateCategoryCenteredModal(props) {
  const [name, setName] = useState("")
  const [rule, setRule] = useState("")
  const initialToken = useSelector(state => state.token)

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const config = {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${initialToken}`
        },
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
      }
      const data = { name, rule }
      const resp = await axios.post(
        'http://localhost:8080/api/v1/category/',
        data, config
      )
      console.log(resp.data)
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Create New Category
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={handleSubmit}>
          <Form.Group controlId="formBasicEmail">
            <Form.Label>Category Name</Form.Label>
            <Form.Control type="text" placeholder="Name" onChange={e => setName(e.target.value)} />
          </Form.Group>
          <Form.Group controlId="exampleForm.ControlTextarea1">
            <Form.Label>Rules</Form.Label>
            <Form.Control as="textarea" rows={3} onChange={e => setRule(e.target.value)} />
          </Form.Group>
          <Button className="btn btn-info btn-block login" type="submit" onSubmit={handleSubmit}>Create</Button>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal >
  );
}

export function CreatePostCenteredModal(props) {
  const [name, setName] = useState("")
  const [rule, setRule] = useState("")
  const initialToken = useSelector(state => state.token)

  const handleSubmit = async e => {
    e.preventDefault();
    try {
      const config = {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${initialToken}`
        },
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
      }
      const data = { name, rule }
      const resp = await axios.post(
        'http://localhost:8080/api/v1/category/',
        data, config
      )
      console.log(resp.data)
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Create New Post
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={handleSubmit}>
          <Form.Group controlId="formBasicEmail">
            <Form.Label>Category Name</Form.Label>
            <Form.Control type="text" placeholder="Name" onChange={e => setName(e.target.value)} />
          </Form.Group>
          <Form.Group controlId="exampleForm.ControlTextarea1">
            <Form.Label>Rules</Form.Label>
            <Form.Control as="textarea" rows={3} onChange={e => setRule(e.target.value)} />
          </Form.Group>
          <Button className="btn btn-info btn-block login" type="submit" onSubmit={handleSubmit}>Create</Button>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal >
  );
}

const Category = (props) => {
  return (
    <li className="nav-item">
      <a className="nav-link" href={props.href} data-toggle="tooltip" data-placement="right" title={props.title}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"
          strokeLinecap="round" strokeLinejoin="round" className="feather feather-dashboard">
          <path d="M3 3 V3 18 C 10 10 20 30 3,3"></path>
        </svg>
        {props.title}
        <span className="sr-only">(current)</span>
      </a>
    </li>
  )
}

export const LeftNav = () => {
  const [modalShow, setModalShow] = useState(false);
  const [component, setComponent] = useState([]);
  const config = {
    headers: {
      'Content-Type': 'application/json'
    },
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  }

  useEffect(() => {
    const initialFetch = async () => {
      const resp = await axios.get(
        'http://localhost:8080/api/v1/category/',
        config
      )
      const newCom = []
      resp.data.forEach((data) => {
        newCom.push(<Category href={"/" + "category" + "/" + data.name} title={data.name} key={data.name}></Category>)
      })
      setComponent(newCom)
    }
    initialFetch()
  }, [])

  return (
    <nav id="leftNav" className="d-md-block sidenav">
      <div className="sidebar-sticky">
        <ul className="nav nav-pill flex-column" style={{ height: "10vh" }}>
        </ul>
        <ul className="nav nav-pill flex-column">
          {component}
        </ul>
        <ul className="nav nav-pill flex-column">
          <li className="nav-item">

          </li>
        </ul>
      </div>
    </nav>
  )
}