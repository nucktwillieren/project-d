import { useState, useContext, createContext } from 'react';
import { useHistory } from 'react-router-dom';
import axios from 'axios';
import { UserContext } from './UserContext';

export const UserContext = createContext(null);

export default function UserPrivateRoute(props) {
  const { user, isLoading } = useContext(UserContext);
  const { component: Component, ...rest } = props;
  if(isLoading) {
     return <Loading/>
  }
  if(user){
     return ( <Route {...rest} render={(props) => 
          (<Component {...props}/>)
           }
        />
      )}
  //redirect if there is no user 
  return <Redirect to='/login' />
}