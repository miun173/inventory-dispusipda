import React from 'react';
import axios from 'axios';
import {
  Table
} from 'antd';

export class Jurnal extends React.Component {
  state = {
    jurnal: [],
  }

  componentDidMount() {
    this.getJurnal();
  }

  getJurnal = async () => {
    const { data } = await axios('/api/jurnal', {
      method: 'GET'
    });

    this.setState({ jurnal: data });
  }

  render() {
    const { jurnal } = this.state;

    const cols = [{
      title: 'Kode',
      dataIndex: 'kode',
      key: 'kode',
    }, {
      title: 'Nama',
      dataIndex: 'nama',
      key: 'nama',
    }, {
      title: 'Reg',
      dataIndex: 'reg',
      key: 'reg',
    }, {
      title: 'Merk',
      dataIndex: 'merk',
      key: 'merk',
    }, {
      title: 'Ukuran',
      dataIndex: 'ukuran',
      key: 'ukuran',
    }, {
      title: 'Bahan',
      dataIndex: 'bahan',
      key: 'bahan',
    }, {
      title: 'Perolehan',
      dataIndex: 'tglMasuk',
      key: 'tglMasuk',
      render: (data) => (
        <span>{(new Date(data)).toLocaleDateString()}</span>
      )
    }, {
      title: 'Tipe Nomor',
      dataIndex: 'tipeSpek',
      key: 'tipeSpek',
    }, {
      title: 'Nomor',
      dataIndex: 'nomorSpek',
      key: 'nomorSpek',
    }]

    return <div style={{ width: '100%', padding: '16px' }}>
      <h2>Jurnal</h2>
      <Table style={{ width: '100%' }} columns={cols} dataSource={jurnal} />
    </div>
  }
}