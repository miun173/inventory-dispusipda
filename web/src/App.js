import React, { Component } from 'react';
import { Icon, Button } from 'antd';
import axios from 'axios';
import styled from 'styled-components';
import {
  BrowserRouter as Router,
  Route,
  Link,
  Switch,
  Redirect,
} from 'react-router-dom';

import {
  BarangMasuk,
  BarangKeluar,
  BukuInventaris,
  Login,
  DivisiRkbmd,
  AccRkbmd,
} from './containers'

import {
  RoutePetugasBarang, RouteDivisi, RoutePejabat,
} from './routes';

import { Provider } from './store';

import { LeftNavComp } from './components';

const Container = styled.div`
  display: flex;
  min-height: 100vh;
`

const initState = {
  userInfo: {
    auth: false,
    role: '',
    id: '',
    name: '',
    token: '',
  },
};

class App extends Component {
  _isMounted = false;
  state = { ...initState }

  componentDidMount() {
    this._isMounted = true;
    this.checkUserInfo();
  }

  componentWillUnmount() {
    this._isMounted = false;
  }

  saveUserInfo = (userInfo) => {
    localStorage.setItem('userInfo', userInfo);
  }

  checkUserInfo = () => {
    const strUserInfo = localStorage.getItem('userInfo');
    if (!strUserInfo) {
      return;
    }
    const userInfo = JSON.parse(strUserInfo);

    // check if auth still valid
    const { id, token, role } = userInfo;
    const creds = btoa(`${id}:${token}`);
    
    axios(`/api/auth/check`, {
      method: 'POST',
      headers: {
      'Authorization': `Basic ${creds}`,
      },
      data: { role }
    })
      .then(() => {
        this.setState({
          userInfo: {
            ...userInfo,
            auth: true,
          },
        })
      })
      .catch((error) => {
        this.setState({
          userInfo: {
            ...initState.userInfo,
            auth: false,
          },
        })
        console.error(error);
      });
  }

  login = async ({ username, password }, cb = () => {}) => {
    try {
      const { data } = await axios('/api/login', {
        method: 'POST',
        data: { username, password }
      });

      if (!this._isMounted) return;
      this.saveUserInfo(JSON.stringify(data));

      this.setState({
        userInfo: { ...data, auth: true },
      }, () => {
        cb(data.role)
      });
    } catch (e) {
      console.error(e);
    }
  }

  logout = () => {
    this.setState({
      ...initState
    });

    return <Redirect to='/' />
  }

  getAuthHeader = () => {
    const { id, token } = this.state.userInfo;
    const creds = btoa(`${id}:${token}`);

    return {'Authorization': `Basic ${creds}`}
  } 

  render() {
    const { userInfo } = this.state;
    return (
      <Provider value={{
        user: userInfo,
        login: this.login,
        authHeader: this.getAuthHeader(),
      }}>
      <Router>
        <Container>
            { userInfo.auth && <LeftNavComp role={userInfo.role} /> }
              <Switch>
                <RoutePetugasBarang path='/inventaris/buku' component={BukuInventaris} />
                <RoutePetugasBarang path='/inventaris/barang-masuk' component={BarangMasuk} />
                <RouteDivisi path='/divisi/inventaris/buku' component={BukuInventaris} />
                <RouteDivisi path='/divisi/barang-keluar' component={BarangKeluar} />
                <RouteDivisi path='/divisi/rkbmd' component={DivisiRkbmd} />
                <RoutePejabat path='/acc/rkbmd' component={AccRkbmd} />

                <Route path='/login' component={Login} />
                <Route path='/logout' component={() => {
                  return this.logout();
                }} />
                <Route path='/' component={() => <main style={{ margin: 'auto' }}>
                  <h1>Home</h1>
                  { !userInfo.auth && <Button><Link to='/login'>Login</Link></Button> }
                </main>} />
              </Switch>
        </Container>
      </Router>
      </Provider>
    );
  }
}

export default App;
