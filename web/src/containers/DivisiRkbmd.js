import React from 'react';
import axios from 'axios';
import styled from 'styled-components';
import {
  Table,
  Button,
  notification,
  Modal,
  Select,
} from 'antd';

import { InputCard } from '../components'

const { Option } = Select

const SidedCard = styled.div`
  display: flex;
`;

const Container = styled.div`
  padding: 8px;
  display: flex;
  /* width: 100%; */
  box-sizing: border-box;
`

const kodeBarang = [
  { kode: '1A1', name: 'Perlengkapan' },
  { kode: '1A2', name: 'Alat tulis kantor' },
  { kode: '1B1', name: 'Kursi' },
  { kode: '1B2', name: 'Meja' },
  { kode: '1B3', name: 'Lemari' },
  { kode: '1B4', name: 'Kendaraan' },
  { kode: '1B5', name: 'Tanah' },
  { kode: '1B6', name: 'Bangunan' },
  { kode: '1B7', name: 'Gedung' },
  { kode: '1B8', name: 'Mesin' },
  { kode: '1C1', name: 'Goodwill' },
  { kode: '1C2', name: 'Paten' },
];
const initState = {
  newBarang: {
    namaBarang: kodeBarang[0].name,
    kodeBarang: kodeBarang[0].kode,
    satuan: '',
    jml: 0,
    harga: 0.0,
  },
  newRkbmd: [],
  rkbmds: [],
  editedBarang: {
    idx: 0,
    namaBarang: '',
    kodeBarang: '',
    satuan: '',
    jml: 0,
    harga: 0.0,
  },
  modVisible: false,
}
export class DivisiRkbmd extends React.Component {
  _isMounted = false
  state = {...initState}
  kodeBarang = kodeBarang
  
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

  handleEditBarangInput = (event) => {
    const target = event.target;
    const value = target.type === 'number' ? parseFloat(target.value) : target.value;
    const name = target.name;

    if (!this._isMounted) return;
    this.setState({
      editedBarang: {
        ...this.state.editedBarang,
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
    try {
      const { data } = await axios('/api/rkbmd', {
        method: 'POST',
        data: {
          rkbmd: this.state.newRkbmd
        }
      });
  
      this.getRkbmd()
  
      this.setState({
        newRkbmd: []
      });

      this.openNotificationWithIcon('success', 'Success Create RKBMD')
    } catch (e) {
      console.error(e)
      this.openNotificationWithIcon('error', 'Failed Create RKBMD')
    }
  }

  openNotificationWithIcon = (type, message, description) => {
    notification[type]({
      message,
      description,
    });
  };

  getRkbmd = async () => {
    const { data } = await axios('/api/rkbmd', {
      method: 'GET'
    });

    this.setState({
      rkbmds: data
    })
  }

  editBarang = (idx) => {
    this.setState({
      modVisible: true,
      editedBarang: {
        idx,
        ...this.state.newRkbmd[idx]
      }
    })
  }

  handleEditNamaBarang = (idx) => {
    if (!this._isMounted) return;
    this.setState({
      editedBarang: {
        ...this.state.editedBarang,
        namaBarang: this.kodeBarang[idx].name,
        kodeBarang: this.kodeBarang[idx].kode,
      }
    });
  }

  handleNamaBarang = (idx) => {
    if (!this._isMounted) return;
    this.setState({
      newBarang: {
        ...this.state.newBarang,
        namaBarang: this.kodeBarang[idx].name,
        kodeBarang: this.kodeBarang[idx].kode,
      }
    });
  }

  handleOk = e => {
    const { newRkbmd, editedBarang } = this.state;

    newRkbmd[editedBarang.idx] = editedBarang

    this.setState({
      modVisible: false,
      newRkbmd,
    });
  };

  handleCancel = e => {
    console.log(e);
    this.setState({
      modVisible: false,
    });
  };

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
      title: 'Harga',
      dataIndex: 'harga',
      key: 'harga',
      render: t => 'Rp' + t.toLocaleString('id-ID')
    }, {
      title: 'Total Harga',
      dataIndex: 'hargaTotal',
      key: 'hargaTotal',
      render: (d, rec) => {
        const sum = rec.harga * rec.jml
        return sum ? 'Rp' + sum.toLocaleString('id-ID') : 'Rp' + 0
      }
    }, {
      title: 'Kode',
      dataIndex: 'kodeBarang',
      key: 'kodeBarang',
    }, {
      title: 'Satuan',
      dataIndex: 'satuan',
      key: 'satuan',
    }, {
      title: 'Action',
      key: 'action',
      render: (_, rec, idx) => {
        return <Button onClick={() => this.editBarang(idx)}>Edit</Button>
      }
    }]
    
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

    const { newBarang, newRkbmd, rkbmds, editedBarang } = this.state

    return <Container>
      <div>
        <h2>Tambah Barang</h2>
        <form>
          <SidedCard>
            <div>
              <label>Nama</label> <br />
              <Select style={{ width: '300px', marginRight: '8px' }} 
                value={newBarang.namaBarang}
                onChange={this.handleNamaBarang}>

                  { this.kodeBarang.map((k, idx) => <Option key={idx} value={idx} >{k.name}</Option>) }
              </Select>
            </div>
            <div>
              <label>Kode</label>
              <p style={{ paddingTop: '4px' }}><b>{newBarang.kodeBarang}</b></p>
            </div>
          </SidedCard>
          <SidedCard>
            <InputCard value={newBarang.satuan} label='Satuan' name='satuan' onChange={this.handleBarangInput}/>
            <InputCard width={100} value={newBarang.jml} type='number' label='Jml' name='jml' onChange={this.handleBarangInput}/>
          </SidedCard>
          <SidedCard>
            <InputCard type='number' value={newBarang.harga} label='Harga' name='harga' onChange={this.handleBarangInput}/>
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

      <Modal title="Edit Barang"
        visible={this.state.modVisible}
        onOk={this.handleOk}
        onCancel={this.handleCancel} >
        
        <form>
          <SidedCard>
            <div>
              <label>Nama</label> <br />
              <Select style={{ width: '300px' }} 
                value={editedBarang.namaBarang}
                name='tipeSpek' onChange={this.handleEditNamaBarang}>

                  { this.kodeBarang.map((k, idx) => <Option key={idx} value={idx} >{k.name}</Option>) }
              </Select>
            </div>
            <div>
              <label>Kode</label>
              <p style={{ paddingTop: '4px' }}><b>{editedBarang.kodeBarang}</b></p>
            </div>
          </SidedCard>
          <SidedCard>
            <InputCard value={editedBarang.satuan} label='Satuan' name='satuan' onChange={this.handleEditBarangInput}/>
            <InputCard width={100} value={editedBarang.jml} type='number' label='Jml' name='jml' onChange={this.handleEditBarangInput}/>
          </SidedCard>
          <SidedCard>
            <InputCard type='number' value={editedBarang.harga} label='Harga' name='harga' onChange={this.handleEditBarangInput}/>
          </SidedCard>
        </form>

        </Modal>
    </Container>
  }
}