import React from 'react';
import axios from 'axios';
import XLSX from 'xlsx';
import { saveAs } from 'file-saver';
import {
  Table,
  Icon,
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

  s2ab = (s) => { 
    const buf = new ArrayBuffer(s.length); //convert s to arrayBuffer
    const view = new Uint8Array(buf);  //create uint8array as viewer
    for (let i=0; i<s.length; i++) view[i] = s.charCodeAt(i) & 0xFF; //convert to octet
    return buf;    
  }

  getExcel = () => {
    const { jurnal } = this.state;
    const data = [];
    const keys = Object.keys(jurnal[0])
    data[0] = ['Id',	'Kode',	'Nama',	'Reg',	'Merk',	'Jml',	'Ukuran',	'Bahan',	'Tipe Nomor',	'Nomor',	'Cara Perolehan',	'Perolehan',	'Harga',	'Nilai Sisa',	'Umur Ekonomis',	'Umur Penggunaan',	'Nilai Buku',	'Beban Penyusutan',	'Ket',	'Harga Total', 'Penyusutan']
    for (let i = 0; i < jurnal.length; i++) {
      const j = jurnal[i]
      let d = []
      for (let k = 0; k < keys.length; k++) {
        console.log(keys[k])
        d.push(j[keys[k]])
      }
      const jdate = (new Date(j.tglMasuk)).toLocaleDateString();
      d[11] = jdate
      data.push(d)
    }

    console.log(data);

    const wb = XLSX.utils.book_new()
    wb.SheetNames.push("Buku Inventaris") 
    wb.Sheets['Buku Inventaris'] = XLSX.utils.aoa_to_sheet(data)
    const wbout = XLSX.write(wb, {bookType:'xlsx',  type: 'binary'});
    saveAs(new Blob([this.s2ab(wbout)],{type:"application/octet-stream"}), 'test.xlsx');
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
      title: 'Harga',
      dataIndex: 'harga',
      key: 'harga',
      render: (d) => d ? 'Rp' + d.toLocaleString('id-ID') : 'Rp' + 0
    }, {
      title: 'Total Harga',
      dataIndex: 'hargaTotal',
      key: 'hargaTotal',
      render: (d) => d ? 'Rp' + d.toLocaleString('id-ID') : 'Rp' + 0
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
      render: (d) => d ? 'Rp' + d.toLocaleString('id-ID') : 'Rp' + 0
    }, {
      title: `Penyusutan ${(new Date()).getFullYear()}`,
      dataIndex: 'penyusutan',
      key: 'penyusutan',
      render: (d) => d ? 'Rp' + d.toLocaleString('id-ID') : 'Rp' + 0
    },]

    return <div style={{ width: '900px', padding: '16px' }}>
      <h2>Buku Inventaris <a onClick={this.getExcel}><Icon type='download' /></a></h2>
      <Table scroll={{ x: 900 }} columns={cols} dataSource={jurnal} />
    </div>
  }
}