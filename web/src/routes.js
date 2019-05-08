import React from 'react';
import { Route, Redirect } from 'react-router-dom';

import { Consumer } from './store'

export const RoutePejabat = ({ component: Component, ...rest }) => (
  <Consumer>
    {({ user }) => {      
      return (
        <Route
          {...rest}
          render={props =>
            user.auth === true && user.role === 'pejabat'
            ? <Component {...props} />
            : <Redirect to="/" />
          }
        />
      )}
    } 
  </Consumer>
);

export const RouteDivisi = ({ component: Component, ...rest }) => (
  <Consumer>
    {({ user }) => {      
      return (
        <Route
          {...rest}
          render={props =>
            user.auth === true && user.role === 'divisi'
            ? <Component {...props} />
            : <Redirect to="/" />
          }
        />
      )}
    } 
  </Consumer>
);

export const RoutePetugasBarang = ({ component: Component, ...rest }) => (
  <Consumer>
    {({ user }) => {      
      return (
        <Route
          {...rest}
          render={props =>
            user.auth === true && user.role === 'petugasBarang'
            ? <Component {...props} />
            : <Redirect to="/" />
          }
        />
      )}
    } 
  </Consumer>
);