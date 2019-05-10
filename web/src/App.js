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
} from './containers'
import {
  RoutePetugasBarang, RouteDivisi,
} from './routes';
import { Provider } from './store';

const Container = styled.div`
  display: flex;
  min-height: 100vh;
`

const LeftNav = styled.div`
  width: 256px;
  background: #222;
`

const LeftMenu = styled.div`
  margin-left: 16px;
  padding: 8px;
  a {
    color: #fff;
  }

  &:hover {
    cursor: pointer;
    opacity: 0.7;
  }
`

const Title = styled.div`
  height: 100px;
  width: 100%;
  background: #1890FF;
  display: flex;
  justify-content: center;
  align-items: center;

  a {
    h3 {
      color: #fff;
      font-weight: bold;
    }
  }
`

const LeftNavComp = ({ role }) => <>
  <LeftNav>
    <Title>
      <Link to='/'>
        <h3>Inventory Barang</h3>
      </Link>
    </Title>
    <br />
    { role === 'petugasBarang' && <>
      <Link to='/inventaris/buku' style={{ color: '#fff' }}>
        <LeftMenu>
          <Icon type='book' style={{ color: '#fff', marginRight: '8px' }} />
          Buku Inventaris
        </LeftMenu>
      </Link>
      <br />

      <Link to='/inventaris/barang-masuk' style={{ color: '#fff' }}>
        <LeftMenu>
          <Icon type='file-search' style={{ color: '#fff', marginRight: '8px' }} />
          Barang Masuk
        </LeftMenu>
      </Link>
      <br />
    </> }

    { role === 'divisi' && <>
      <Link to='/barang-keluar' style={{ color: '#fff' }}>
        <LeftMenu>
        <Icon type='export' style={{ color: '#fff', marginRight: '8px' }} />
        Barang Keluar
        </LeftMenu>
      </Link>
      <br />
    </> }

    { role === 'pejabat' && <>
      <Link to='/acc/rkbmd' style={{ color: '#fff' }}>
        <LeftMenu>
          <Icon type='file-protect' style={{ color: '#fff', marginRight: '8px' }} />
          RKBMD
        </LeftMenu>
      </Link>
      <br />
    </> }

    <Link to='/logout' style={{ color: '#fff' }}>
      <LeftMenu>
          <Icon type="logout" style={{ color: '#fff', marginRight: '8px' }} />
          Logout
      </LeftMenu>
    </Link>
    <br />

  </LeftNav>
</>

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
  }

  componentWillUnmount() {
    this._isMounted = false;
  }

  login = async ({ username, password }, cb = () => {}) => {
    try {
      const { data } = await axios('/api/login', {
        method: 'POST',
        data: { username, password }
      });

      console.log(data);
      if (!this._isMounted) return;
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

  getAuthHeader = () => this.state.userInfo.token 

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
                <RouteDivisi path='/barang-keluar' component={BarangKeluar} />
                <Route path='/login' component={Login} />
                <Route path='/logout' component={() => {
                  return this.logout();
                }} />
                <Route path='/' component={() => <main>
                  <h1>Home</h1>
                  <Button><Link to='/login'>Login</Link></Button>
                </main>} />
              </Switch>
        </Container>
      </Router>
      </Provider>
    );
  }
}

export default App;
