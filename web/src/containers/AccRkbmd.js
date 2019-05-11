import React from 'react';
import axios from 'axios';
import styled from 'styled-components';

import { Table, Button } from 'antd';

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
  },
  rkbdmdSlectedRowKeys: [],
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
    const { selectedRkbmds, rkbdmdSlectedRowKeys: rks } = this.state
    const rkDetail = selectedRkbmds
      .details.map(r => ({ ...r, status: 'decline' }))

    rks.forEach((r) => {
      rkDetail[r].status = 'acc'
    });

    await axios('/api/rkbmd', {
      method: 'PUT',
      data: {
        ...this.state.selectedRkbmds,
        details: rkDetail,
      }
    });

    await this.getRkbmds()
    this.setState({
      selectedRkbmds: this.state.rkbmds.find(r => r.id === selectedRkbmds.id)
    });
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
        <h2>Detail RKBMD {selectedRkbmds.id}</h2>
        <Table rowSelection={rkbmdItemSelection} columns={detailRkbmd} dataSource={selectedRkbmds.details} />
        <Button style={{ width: '100%' }} type='primary' onClick={this.onAccRkbmd}>ACC</Button>
      </div>}
    </Container>
  }
}