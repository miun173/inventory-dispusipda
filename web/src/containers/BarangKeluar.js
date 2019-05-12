import React from 'react';
import styled from 'styled-components';
import axios from 'axios';
import XLSX from 'xlsx';
import { saveAs } from 'file-saver';

import { Input, Button, notification, Icon, Table } from 'antd';

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
      this.getBarangKel();
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

  s2ab = (s) => { 
    const buf = new ArrayBuffer(s.length); //convert s to arrayBuffer
    const view = new Uint8Array(buf);  //create uint8array as viewer
    for (let i=0; i<s.length; i++) view[i] = s.charCodeAt(i) & 0xFF; //convert to octet
    return buf;    
  }

  getExcel = () => {
    const { barangs } = this.state;
    const data = [];
    const keys = Object.keys(barangs[0])

    // Insert RKBMD to rows
    data[0] = ['ID', 'Barang ID', 'Nama', 'Jumlah', 'Tgl Keluar']

    for (let i = 0; i < barangs.length; i++) {
      const j = barangs[i];
      const jdate = (new Date(j.tglKeluar)).toLocaleDateString();
      data.push([j.id, j.barangID, j.nama, j.jml, jdate])
    }

    console.log(data)

    const wb = XLSX.utils.book_new()
    wb.SheetNames.push('Barang Keluar')
    wb.Sheets['Barang Keluar'] = XLSX.utils.aoa_to_sheet(data)

    const wbout = XLSX.write(wb, {bookType:'xlsx',  type: 'binary'});
    saveAs(new Blob([this.s2ab(wbout)],{type:"application/octet-stream"}), 'test.xlsx');
  }

  render() {
    const { newBarangKel, barangs } = this.state;

    const cols = [{
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    }, {
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
      <h2>Barang Keluar <a onClick={this.getExcel}><Icon type='download' /></a></h2>
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