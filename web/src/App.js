import React, { Component } from 'react';
import { Icon } from 'antd';
import axios from 'axios';
import styled from 'styled-components';
import {
  BrowserRouter as Router,
  Route,
  Link,
  Switch,
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
  a {
    color: #fff;
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

const LeftNavComp = () => <>
  <LeftNav>
    <Title>
      <Link to='/'>
        <h3>Inventory Barang</h3>
      </Link>
    </Title>
    <br />
    <LeftMenu>
      <Icon type='book' style={{ color: '#fff', marginRight: '8px' }} />
      <Link to='/inventaris/buku'>Buku Inventaris</Link>
    </LeftMenu>
    <br />
    <LeftMenu>
      <Icon type='import' style={{ color: '#fff', marginRight: '8px' }} />
      <Link to='/inventaris/barang-masuk'>Barang Masuk</Link>
    </LeftMenu>
    <br />
    <LeftMenu>
      <Icon type='export' style={{ color: '#fff', marginRight: '8px' }} />
      <Link to='/barang-keluar'>Barang Keluar</Link>
    </LeftMenu>
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
      if (this._isMounted) return;
      this.setState({
        userInfo: { ...data, auth: true },
      }, () => {
        cb();
      });
    } catch (e) {
      console.error(e);
    }
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
            { userInfo.auth && <LeftNavComp /> }
              <RoutePetugasBarang path='/inventaris/barang-masuk' component={BarangMasuk} />
              <RoutePetugasBarang path='/inventaris/buku' component={BukuInventaris} />
              <RouteDivisi path='/barang-keluar' component={BarangKeluar} />
              <Route path='/' component={Login} />
        </Container>
      </Router>
      </Provider>
    );
  }
}

export default App;
