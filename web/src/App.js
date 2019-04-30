import React, { Component } from 'react';
import styled from 'styled-components';
import {
  BrowserRouter as Router,
  Route,
  Link,
} from 'react-router-dom';
import {
  BarangMasuk,
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
`

const Title = styled.div`
  height: 100px;
  width: 100%;
  background: blue;
  display: flex;
  justify-content: center;
  align-items: center;
`

class App extends Component {
  render() {
    return (
      <Router>
      <Container>
        <LeftNav>
          <Title>
            <Link to='/'>Inventory Barang</Link>
          </Title>
          <br />
          <LeftMenu>
            <Link to='/barang'>Barang Masuk</Link>
          </LeftMenu>
        </LeftNav>
          <Route path='/barang' component={BarangMasuk} />
      </Container>
      </Router>
    );
  }
}

export default App;
