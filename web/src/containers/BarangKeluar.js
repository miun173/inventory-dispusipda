import React from 'react';
import styled from 'styled-components';
import axios from 'axios';

import { Input, Button, notification, Table } from 'antd';

const Container = styled.div`
  padding: 16px;
  box-sizing: border-box;
`

const Card = styled.div`
  padding: 8px;
  width: 300px;
`

const initState = {
  newBarangKel: {
    barangID: 0,
    jml: 0,
    tglKeluar: null,
  },
  barangs: [],
};

export class BarangKeluar extends React.Component {
  _isMounted = false
  state = {
    ...initState
  }

  componentDidMount() {
    this._isMounted = true
    this.getBarangKel();
  }

  componentWillUnmount() {
    this._isMounted = false
  }

  getBarangKel = async () => {
    const { data } = await axios('/api/barang-keluar', {
      method: 'GET',
    });

    this.setState({
      barangs: data,
    })
  } 

  handleBarangInput = (event) => {
    const target = event.target;
    const value = target.type === 'number' ? parseFloat(target.value) : target.value;
    const name = target.name;

    if (!this._isMounted) return;
    this.setState({
      newBarangKel: {
        ...this.state.newBarangKel,
        [name]: value,
      }
    });
  }

  handleSubmit = async (e) => {
    e.preventDefault();
    const { newBarangKel } = this.state;

    try {
      const { data } = await axios('/api/barang-keluar', {
        method: 'POST',
        data: {
          ...newBarangKel,
          tglKeluar: (new Date(newBarangKel.tglKeluar)).getTime(),
        }
      });

      this.setState({ newBarangKel: initState.newBarangKel })
      this.openNotificationWithIcon('success', 'Success');
      console.log(data);
    } catch (e) {
      this.openNotificationWithIcon('error', 'Field cannot be empty');
    }
  };

  openNotificationWithIcon = (type, message, description) => {
    notification[type]({
      message,
      description,
    });
  };

  render() {
    const { newBarangKel, barangs } = this.state;

    const cols = [{
      title: 'ID Barang',
      dataIndex: 'barangID',
      key: 'barangId',
    }, {
      title: 'Nama',
      dataIndex: 'nama',
      key: 'nama',
    }, {
      title: 'Jumlah',
      dataIndex: 'jml',
      key: 'jml',
    }, {
      title: 'Tgl Keluar',
      dataIndex: 'tglKeluar',
      key: 'tglKeluar',
      render: (d) => <span>{(new Date(d)).toLocaleDateString()}</span> 
    }, ];

    return <Container>
      <h2>Barang Keluar</h2>
      <form style={{ display: 'flex', }}>
        <Card style={{ width: '100px' }}>
            <label>ID Barang</label> <br />
            <Input type='number' style={{ width: '80px' }} value={newBarangKel.barangID} name='barangID' onChange={this.handleBarangInput}/>
        </Card>
        <Card>
            <label>Nama</label>
            <Input type='text' value={newBarangKel.nama} name='nama' onChange={this.handleBarangInput}/>
        </Card>
        <Card style={{ width: '100px' }}>
            <label>Jumlah</label> <br />
            <Input  type='number' value={newBarangKel.jml} name='jml' onChange={this.handleBarangInput}/>
        </Card>
        <Card style={{ width: '180px' }}>
            <label>Tgl Keluar</label> <br />
            <Input type='date' value={newBarangKel.tglKeluar} name='tglKeluar' onChange={this.handleBarangInput}/>
        </Card>

        <Card>
          <label></label> <br />
          <Button
            onClick={this.handleSubmit}
            style={{ width: '100px' }} 
            htmlType='submit' type='primary'>
            Tambah
          </Button>
        </Card>
      </form>

      <Table style={{ width: '800px' }} columns={cols} dataSource={barangs} scroll={{ x: 800 }} />
    </Container>
  }
}