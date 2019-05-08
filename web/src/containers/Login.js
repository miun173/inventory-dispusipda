import React from 'react';

import { Card, Input, Button, Icon } from 'antd';
import { Redirect, withRouter } from 'react-router-dom';
import { AppContext } from '../store'

const LoginForm = ({ username, password, onLogin, handlePoInput, title }) => (
  <Card style={{ paddingBottom: '16px', maxWidth: '300px', margin: 'auto', boxSizing: 'border-box' }}>
    <form onSubmit={onLogin}>
      <h1>{title || 'Login'}</h1>
      <div style={{ marginTop: '8px', width: '100%', paddingTop: '0px', boxSizing: 'border-box' }}>
        <label> Username </label> <br />
        <Input value={username}
        placeholder='Username'
        style={{ width: '100%' }}
        prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}  
        name='username' onChange={handlePoInput}/>
      </div>
  
      <div style={{ marginTop: '8px', width: '100%', boxSizing: 'border-box' }}>
        <label> Password </label> <br />
        <Input value={password} 
          placeholder='Password'
          type='password'
          style={{ width: '100%' }}
          prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
          name='password' onChange={handlePoInput}/>
      </div>
      <Button
        type='primary'
        htmlType='submit'
        onClick={onLogin}
        style={{ marginTop: '16px', width: '100%', boxSizing: 'border-box' }} 
      >Login</Button>
    </form>
  </Card>
);

class LoginComp extends React.Component {
  static contextType = AppContext;

  state = {
    username: '',
    password: '',
  }

  handlePoInput = (event) => {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  render() {
    const { username, password } = this.state;

    return this.context.user.auth 
      ? <Redirect to='/store/manager' /> 
      : (<div style={{ marginTop: '32px', width: '100%' }}>
        <LoginForm 
              title='Login Store'
              username={username} 
              password={password} 
              handlePoInput={this.handlePoInput}
              onLogin={(e) => {
                e.preventDefault();
                this.context.login({ username, password }, () => {
                  this.props.history.push('/store/manager');
                })} 
              } />
      </div>);
  }
};

export const Login = withRouter(LoginComp);