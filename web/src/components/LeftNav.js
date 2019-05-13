import React from 'react';
import styled from 'styled-components';

import { Icon } from 'antd';
import { Link } from 'react-router-dom';

export const LeftNav = styled.div`
  width: 256px;
  background: #222;
`

export const LeftMenu = styled.div`
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

export const Title = styled.div`
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

const user = {
  pejabat: 'Subag Umum',
  petugasBarang: 'Petugas Barang',
  divisi: 'Divisi',
}

export const LeftNavComp = ({ role }) => <>
<LeftNav>
  <Title>
    <Link style={{ textAlign: 'center' }} to='/'>
      <h3>Inventory Barang</h3>
      <h4 style={{ color: '#fff' }}>{ user[role] }</h4>
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

    {/* <Link to='/inventaris/barang-keluar' style={{ color: '#fff' }}>
      <LeftMenu>
      <Icon type='export' style={{ color: '#fff', marginRight: '8px' }} />
      Barang Keluar
      </LeftMenu>
    </Link>
    <br /> */}

    <Link to='/inventaris/barang-masuk' style={{ color: '#fff' }}>
      <LeftMenu>
        <Icon type='file-search' style={{ color: '#fff', marginRight: '8px' }} />
        Barang Masuk
      </LeftMenu>
    </Link>
    <br />
  </> }

  { role === 'divisi' && <>
    <Link to='/divisi/inventaris/buku' style={{ color: '#fff' }}>
      <LeftMenu>
      <Icon type='book' style={{ color: '#fff', marginRight: '8px' }} />
      Buku Inventaris
      </LeftMenu>
    </Link>
    <br />
    {/* <Link to='/divisi/barang-keluar' style={{ color: '#fff' }}>
      <LeftMenu>
      <Icon type='export' style={{ color: '#fff', marginRight: '8px' }} />
      Barang Keluar
      </LeftMenu>
    </Link>
    <br /> */}
    <Link to='/divisi/rkbmd' style={{ color: '#fff' }}>
      <LeftMenu>
      <Icon type='create' style={{ color: '#fff', marginRight: '8px' }} />
      Buat RKBMD
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

