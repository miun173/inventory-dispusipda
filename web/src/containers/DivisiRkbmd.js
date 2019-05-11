import React from 'react';
import axios from 'axios';
import styled from 'styled-components';
import {
  Table,
  Button,
} from 'antd';

import { InputCard } from '../components'

const SidedCard = styled.div`
  display: flex;
`;

const Container = styled.div`
  padding: 8px;
  display: flex;
  /* width: 100%; */
  box-sizing: border-box;
`

const initState = {
  newBarang: {
    namaBarang: '',
    kodeBarang: '',
    satuan: '',
    jml: 0,
  },
  newRkbmd: [],
  rkbmds: []
}
export class DivisiRkbmd extends React.Component {
  _isMounted = false
  state = {...initState}
  
  componentWillUnmount() {
    this._isMounted = false
  }

  componentDidMount() {
    this._isMounted = true
    this.getRkbmd()
  }

  handleBarangInput = (event) => {
    const target = event.target;
    const value = target.type === 'number' ? parseFloat(target.value) : target.value;
    const name = target.name;

    if (!this._isMounted) return;
    this.setState({
      newBarang: {
        ...this.state.newBarang,
        [name]: value,
      }
    });
  }

  handleTambah = (e) => {
    e.preventDefault();
    if (!this._isMounted) return
    this.setState({
      newRkbmd: [
        ...this.state.newRkbmd,
        {...this.state.newBarang}
      ],
      newBarang: initState.newBarang,
    })
  }

  createRkbmd = async () => {
    const { data } = await axios('/api/rkbmd', {
      method: 'POST',
      data: {
        rkbmd: this.state.newRkbmd
      }
    });

    this.getRkbmd()

    this.setState({
      newRkbmd: []
    })
  }

  getRkbmd = async () => {
    const { data } = await axios('/api/rkbmd', {
      method: 'GET'
    });

    this.setState({
      rkbmds: data
    })
  }

  render() {
    const tableBarang = [{
      title: 'Nama',
      dataIndex: 'namaBarang',
      key: 'nama',
    }, {
      title: 'Jumlah',
      dataIndex: 'jml',
      key: 'jumlah',
    }, {
      title: 'Kode',
      dataIndex: 'kode',
      key: 'kode',
    }, {
      title: 'Satuan',
      dataIndex: 'satuan',
      key: 'satuan',
    },]
    
    const tableRkbmd = [{
      title: 'ID',
      dataIndex: 'id',
      key: 'id'
    }, {
      title: 'Tgl Diajukan',
      dataIndex: 'tglBuat',
      key: 'tglBuat',
      render: d => (new Date(d)).toLocaleDateString(),
    }, {
      title: 'Status',
      dataIndex: 'status',
      key: 'status'
    }]

    const { newBarang, newRkbmd, rkbmds } = this.state
    return <Container>
      <div>
        <h2>Tambah Barang</h2>
        <form>
          <SidedCard>
            <InputCard value={newBarang.namaBarang} label='Nama Barang' name='namaBarang' onChange={this.handleBarangInput}/>
            <InputCard width={100} value={newBarang.jml} type='number' label='Jml' name='jml' onChange={this.handleBarangInput}/>
          </SidedCard>
          <SidedCard>
            <InputCard value={newBarang.satuan} label='Satuan' name='satuan' onChange={this.handleBarangInput}/>
            <InputCard width={100} value={newBarang.kodeBarang} label='Kode Barang' name='kodeBarang' onChange={this.handleBarangInput}/>
          </SidedCard>
          <Button htmlType='submit' onClick={this.handleTambah} style={{ width: '100%' }} type='primary'>Tambah</Button>
        </form>
        <br />

        <h2>Daftar Barang Yang Diajukan</h2>
        <Table columns={tableBarang} dataSource={newRkbmd} />
        <Button style={{ width: '100%' }} onClick={this.createRkbmd} type='primary'>Buat RKBMD</Button>
      </div>

      <div style={{ marginLeft: '64px' }}>
        <h2>Daftar RKBMD</h2>
        <Table columns={tableRkbmd} dataSource={rkbmds} />
      </div>
    </Container>
  }
}