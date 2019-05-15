import React from 'react';
import axios from 'axios';
import styled from 'styled-components';
import { saveAs } from 'file-saver';
import XLSX from 'xlsx';

import { Table, Button, notification, Icon } from 'antd';

const Container = styled.div`
  display: flex;
  flex-direction: row;
  padding: 8px;
  width: 100%;
  justify-content: flex-start;
`

const initState = {
  rkbmds: [],
  selectedRkbmds: {
    id: null,
    details: [],
    tglBuat: null,
    status: 'acc',
  },
  rkbdmdSlectedRowKeys: [],
  isAllDecline: false,
};

export class AccRkbmd extends React.Component {
  _isMounted = false
  state = {...initState}

  componentWillUnmount() {
    this._isMounted = false;
  }

  componentDidMount() {
    this._isMounted = true
    this.getRkbmds();
  }

  openNotificationWithIcon = (type, message, description) => {
    notification[type]({
      message,
      description,
    });
  };

  getRkbmds = () => new Promise(async (resolve, reject) => {
    try {
      const { data } = await axios('/api/rkbmd', {
        method: 'GET',
      });
  
      console.log(data);
  
      this.setState({ rkbmds: data }, () => {
        resolve();
      });
    } catch (e) {
      reject(e);
    }
  });

  handleDetail = (rkbmdID) => {
    this.setState({
      selectedRkbmds: this.state.rkbmds.find(r => r.id === rkbmdID)
    })
  }

  onAccRkbmd = async () => {
    try {
      const { selectedRkbmds, rkbdmdSlectedRowKeys: rks } = this.state
      const rkDetail = selectedRkbmds
        .details.map(r => ({ ...r, status: 'decline' }))
  
      rks.forEach((r) => {
        rkDetail[r].status = 'acc'
      });

      // if all detail arr declined, rk is declined
      const isAllDecline = rkDetail.every(r => r.status === 'decline' || r.status === 'pending')
      console.log('isAllDecline', isAllDecline)
      selectedRkbmds.status = isAllDecline ? 'decline' : 'acc'
  
      await axios('/api/rkbmd', {
        method: 'PUT',
        data: {
          ...selectedRkbmds,
          details: rkDetail,
        }
      });
  
      await this.getRkbmds()
      this.setState({
        selectedRkbmds: this.state.rkbmds.find(r => r.id === selectedRkbmds.id)
      });
  
      this.openNotificationWithIcon('success', 'Success ACC');
    } catch (e) {
      this.openNotificationWithIcon('error', 'Failed to ACC');
    }
  }

  onSelectChange = (selectedRowKeys) => {
    console.log(selectedRowKeys)
    const rk = this.state.selectedRkbmds
    for (let i = 0; i < rk.details.length; i++) {
      // if (rk.details[i].status !== 'acc') {
        rk.details[i].status = 'decline'
      // }
    }

    for (let i = 0; i < selectedRowKeys.length; i++) {
      rk.details[selectedRowKeys[i]].status = 'acc'
    }

    this.setState({
      rkbdmdSlectedRowKeys: selectedRowKeys,
      selectedRkbmds: rk
    });
  }

  s2ab = (s) => { 
    const buf = new ArrayBuffer(s.length); //convert s to arrayBuffer
    const view = new Uint8Array(buf);  //create uint8array as viewer
    for (let i=0; i<s.length; i++) view[i] = s.charCodeAt(i) & 0xFF; //convert to octet
    return buf;    
  }

  getExcel = () => {
    const { selectedRkbmds: srk } = this.state;
    const rkDetail = srk.details;
    const data = [];
    const keys = Object.keys(rkDetail[0])

    // Insert RKBMD to rows
    data[0] = ['Rkbmd ID', 'Tgl Buat', 'Status']
    const sDate = (new Date(srk.tglBuat)).toLocaleDateString()
    data[1] = [srk.id, sDate, srk.status]
    data[2] = []


    // insert rkDetail to rows
    data[3] = ['ID','Rkbmd ID','Jumlah','Kode Barang', 'Nama Barang', 'Status', 'Harga', 'Total Harga']
    for (let i = 0; i < rkDetail.length; i++) {
      const j = rkDetail[i];
      let d = [];

      for (let k = 0; k < keys.length; k++) {
        console.log(keys[k])
        d.push(j[keys[k]])
      }

      const jdate = (new Date(j.tglMasuk)).toLocaleDateString();
      d['tglBuat'] = jdate
      data.push(d)
    }

    console.log(data)

    const wb = XLSX.utils.book_new()
    wb.SheetNames.push('RKBMD')
    wb.Sheets['RKBMD'] = XLSX.utils.aoa_to_sheet(data)

    const wbout = XLSX.write(wb, {bookType:'xlsx',  type: 'binary'});
    saveAs(new Blob([this.s2ab(wbout)],{type:"application/octet-stream"}), 'test.xlsx');
  }

  render() {
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
    }, {
      title: 'Action',
      key: 'action',
      render: (_, r) => <Button onClick={() => this.handleDetail(r.id)}>Detail</Button>
    }];

    const detailRkbmd = [{
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    }, {
      title: 'Nama Barang',
      dataIndex: 'namaBarang',
      key: 'namaBarang'
    }, {
      title: 'Jumlah',
      dataIndex: 'jml',
      key: 'jml',
    }, {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    }]

    const { rkbmds, selectedRkbmds, rkbdmdSlectedRowKeys } = this.state;

    const rkbmdItemSelection = {
      rkbdmdSlectedRowKeys,
      onChange: this.onSelectChange,
    };

    return <Container>
      <div>
        <h2>RKBMD</h2>
        <Table columns={tableRkbmd} dataSource={rkbmds} />
      </div>

      {selectedRkbmds.id && <div style={{ padding: '0 16px' }}>
        <h2>Detail RKBMD {selectedRkbmds.id} <a onClick={this.getExcel}><Icon type='download' /></a></h2>
        <Table rowSelection={rkbmdItemSelection} columns={detailRkbmd} dataSource={selectedRkbmds.details} />
        <Button style={{ width: '100%' }} type='primary' onClick={this.onAccRkbmd}>ACC</Button>
      </div>}
    </Container>
  }
}