import React from 'react';
import axios from 'axios';
import {
  Table
} from 'antd';

export class BukuInventaris extends React.Component {
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
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    }, {
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
      title: 'Jml',
      dataIndex: 'jml',
      key: 'jml',
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
    }, {
      title: 'Keterangan',
      dataIndex: 'ket',
      key: 'ket',
    }, {
      title: 'Harga',
      dataIndex: 'harga',
      key: 'harga',
      render: (d) => d ? d.toLocaleString('id-ID') : 0
    }, {
      title: 'Nilai Sisa',
      dataIndex: 'nilaiSisa',
      key: 'nilaiSisa',
      render: (d) => d ? d.toLocaleString('id-ID') : 0
    }, {
      title: 'Umur Ekonomis',
      dataIndex: 'umurEkonomis',
      key: 'umurEkonomis',
    }, {
      title: 'Umur Penggunaan',
      dataIndex: 'umurPenggunaan',
      key: 'umurPenggunaan',
    }, {
      title: 'Nilai Buku',
      dataIndex: 'nilaiBuku',
      key: 'nilaiBuku',
      render: (d) => d ? d.toLocaleString('id-ID') : 0
    }, {
      title: 'Beban Penyusutan',
      dataIndex: 'bebanPenyusutan',
      key: 'bebanPenyusutan',
      render: (d) => d ? d.toLocaleString('id-ID') : 0
    }, {
      title: `Penyusutan ${(new Date()).getFullYear()}`,
      dataIndex: 'penyusutan',
      key: 'penyusutan',
      render: (d) => d ? d.toLocaleString('id-ID') : 0
    },]

    return <div style={{ width: '900px', padding: '16px' }}>
      <h2>Jurnal</h2>
      <Table scroll={{ x: 900 }} columns={cols} dataSource={jurnal} />
    </div>
  }
}