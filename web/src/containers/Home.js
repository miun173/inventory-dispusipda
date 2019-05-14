import React from 'react';
import { Button } from 'antd';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

import ImgMngmnt from '../icons/management.svg'

const Nav = styled.nav`
  display: flex;
  width: 100%;
  background: #222;
  justify-content: space-between;
  padding: 8px 32px;

  h1 {
    color: #fff;
  }
`

const Main = styled.main`
  display: flex;
  padding: 16px;
  /* padding-left: 64px; */
  justify-content: center;
  align-items: center;
`

const ImgHome = styled.div`
  background: ${({ url }) => `url(${url})`};
  background-size: cover;
  width: 450px;
  height: 350px;
`

export const Home = ({ userInfo }) => {
  return <>
  <div style={{ width: '100%' }}>
    <Nav>
    <h1 style={{ margin: '0 8px' }}>Home</h1>
      { !userInfo.auth && <Button><Link to='/login'>Login</Link></Button> }      
    </Nav>
    <Main>
      <div>
        <h1 style={{ width: '300px' }}>Welcome to Assets Management Cycle</h1>
        <h2>Dispusipda Jawa Barat</h2>
        <h3 style={{ width: '400px' }}>Perpustakaan Mencerdaskan Masyarakat dan Kearsipan Pilar Akuntabilitas</h3>
      </div>

      <ImgHome url={ImgMngmnt} />
    </Main>
  </div>
  </>
}