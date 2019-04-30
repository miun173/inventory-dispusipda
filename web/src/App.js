import React, { Component } from 'react';
import { Icon } from 'antd';
import styled from 'styled-components';
import {
  BrowserRouter as Router,
  Route,
  Link,
} from 'react-router-dom';
import {
  BarangMasuk,
  Jurnal,
} from './containers'

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

class App extends Component {
  render() {
    return (
      <Router>
      <Container>
        <LeftNav>
          <Title>
            <Link to='/'>
              <h3>Inventory Barang</h3>
            </Link>
          </Title>
          <br />
          <LeftMenu>
            <Icon type='book' style={{ color: '#fff', marginRight: '8px' }} />
            <Link to='/jurnal'>Jurnal</Link>
          </LeftMenu>
          <br />
          <LeftMenu>
            <Icon type='export' style={{ color: '#fff', marginRight: '8px' }} />
            <Link to='/barang'>Barang Masuk</Link>
          </LeftMenu>
          <br />
          <LeftMenu>
            <Icon type='import' style={{ color: '#fff', marginRight: '8px' }} />
            <Link to='/barang-keluar'>Barang Keluar</Link>
          </LeftMenu>
        </LeftNav>
          <Route path='/barang' component={BarangMasuk} />
          <Route path='/jurnal' component={Jurnal} />
      </Container>
      </Router>
    );
  }
}

export default App;
