import React from 'react';
import axios from 'axios';
import {
  Input,
  Select,
  Button,
  notification,
  Table,
  Modal,
} from 'antd';
import styled from 'styled-components';

const { Option } = Select;

const Container = styled.div`
  padding: 16px;
  box-sizing: border-box;
  display: flex;
  justify-content: space-between;
`

const Card = styled.div`
  padding: 8px;
  width: 300px;
`

const SidedCard = styled.div`
  display: flex;
`

const tipe = ['Pabrik', 'Polisi', 'Rangka', 'Mesin', 'BPKB']
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

const initNewBarang = {
  kode: kodeBarang[0].kode, 
  nama: kodeBarang[0].name, 
  reg: '',
  merk: '',
  ukuran: '',
  bahan: '',
  tglMasuk: (new Date()).getTime(),
  tipeSpek: tipe[0],
  nomorSpek: '',
  caraPerolehan: '',
  jml: 0,
  harga: 0.0,
  ket: '',
  nilaiSisa: 0,
  umurEkonomis: 0,
  umurPenggunaan: 0,
  nilaiBuku: 0,
  bebanPenyusutan: 0.0,
  koreksi: 0.0,
}

export class BarangMasuk extends React.Component {
  _isMounted = false
  tipe = tipe
  kodeBarang = kodeBarang
  state = {
    barangs: [],
    modVisible: false,
    newBarang: { ...initNewBarang },
    editedBarang: {
      idx: 0,
      id: 0,
      ...initNewBarang
    },
  }

  componentDidMount() {
    this._isMounted = true
    this.getBarangs();
  }

  openNotificationWithIcon = (type, message, description) => {
    notification[type]({
      message,
      description,
    });
  };

  handleSelectSpek = (value) => {
    if (!this._isMounted) return;
    this.setState({
      newBarang: {
        ...this.state.newBarang,
        tipeSpek: value,
      }
    });
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

  handleSubmit = async (e) => {
    e.preventDefault();
    const { newBarang } = this.state;

    try {
      const { data } = await axios('/api/barang', {
        method: 'POST',
        data: {
          ...newBarang,
          tglMasuk: (new Date(newBarang.tglMasuk)).getTime(),
        }
      });
      this.setState({ newBarang: initNewBarang })
      this.openNotificationWithIcon('success', 'Success add barang');
      this.getBarangs();
      console.log(data);
    } catch (e) {
      this.openNotificationWithIcon('error', 'Failed add barang');
    }
  };

  getBarangs = async () => {
    try {
      const { data } = await axios('/api/barang', {
        method: 'GET'
      })

      if (this._isMounted) {
        this.setState({ barangs: data })
      }
    } catch (e) {
      this.openNotificationWithIcon('error', 'cannot list barang')
      console.error(e);
    }
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

  editBarangMasuk = (idx) => {
    this.setState({
      modVisible: true,
      editedBarang: {
        idx,
        ...this.state.barangs[idx]
      }
    })
  }

  handleEditSpek = (value) => {
    if (!this._isMounted) return;
    this.setState({
      editedBarang: {
        ...this.state.editedBarang,
        tipeSpek: value,
      }
    });
  }

  handleEditNamaBarang = (idx) => {
    if (!this._isMounted) return;
    this.setState({
      editedBarang: {
        ...this.state.editedBarang,
        nama: this.kodeBarang[idx].name,
        kode: this.kodeBarang[idx].kode,
      }
    });
  }

  handleNamaBarang = (value) => {
    if (!this._isMounted) return;
    this.setState({
      newBarang: {
        ...this.state.newBarang,
        nama: this.kodeBarang[value].name,
        kode: this.kodeBarang[value].kode
      }
    });
  }

  handleOk = async (e) => {
    e.preventDefault();
    const { editedBarang } = this.state;
    try {
      await axios('/api/barang', {
        method: 'PUT',
        data: {
          ...editedBarang,
          tglMasuk: (new Date(editedBarang.tglMasuk)).getTime(),
        }
      })
  
      this.setState({
        modVisible: false,
      });

      this.getBarangs();

      this.openNotificationWithIcon('success', 'Success update barang');
    } catch (e) {
      console.error(e);
      this.openNotificationWithIcon('error', 'Failed update barang');      
    }
  };

  handleCancel = e => {
    this.setState({
      modVisible: false,
    });
  };

  render() {
    const { newBarang, barangs, editedBarang } = this.state;

    const cols = [{
      title: 'ID Barang',
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
      title: 'Jumlah',
      dataIndex: 'jml',
      key: 'jml',
    }, {
      title: 'Tgl Masuk',
      dataIndex: 'tglMasuk',
      key: 'tglMasuk',
      render: (d) => <span>{(new Date(d)).toLocaleDateString()}</span> 
    }, {
      title: 'Action',
      key: 'action',
      render: (_, rec, idx) => {
        return <Button onClick={() => this.editBarangMasuk(idx)}>Edit</Button>
      }
    }];

    const editBarangDate = new Date(this.state.editedBarang.tglMasuk)
    console.log('editBarangDate', editBarangDate)

    return <Container>
      <div>
        <h2>Barang Masuk</h2>
        <form>
          <SidedCard>
            <Card>
              <label>Nama</label> <br />
              <Select style={{ width: '100%' }} 
                value={newBarang.nama}
                name='tipeSpek' onChange={this.handleNamaBarang}>

                  { this.kodeBarang.map((k, idx) => <Option key={idx} value={idx} >{k.name}</Option>) }
              </Select>
            </Card>
            <Card>
              <label>Kode</label> <br />
              <p style={{ paddingTop: '4px' }}><b>{newBarang.kode}</b></p>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>No.Reg</label>
              <Input value={newBarang.reg} name='reg' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Merk</label>
              <Input value={newBarang.merk} name='merk' onChange={this.handleBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Ukuran/CC</label>
              <Input value={newBarang.ukuran} name='ukuran' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Bahan</label>
              <Input value={newBarang.bahan} name='bahan' onChange={this.handleBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Perolehan</label>
              <Input type='date' value={newBarang.tglMasuk} name='tglMasuk' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Nomor</label>
              <div style={{
                display: 'flex',
              }}>
              <Select style={{ width: '100px' }} 
                value={newBarang.tipeSpek}
                name='tipeSpek' onChange={this.handleSelectSpek}>

                  { this.tipe.map((t, idx) => <Option key={idx} value={t} >{t}</Option>) }
              </Select>
              <Input style={{ marginLeft: '8px' }} value={newBarang.nomorSpek} name='nomorSpek' onChange={this.handleBarangInput}/>
              </div>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Cara Perolehan/Status</label>
              <Input value={newBarang.caraPerolehan} name='caraPerolehan' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Harga</label>
              <Input type='number' value={newBarang.harga} name='harga' onChange={this.handleBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Jumlah</label>
              <Input type='number' value={newBarang.jml} name='jml' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Keterangan</label>
              <Input value={newBarang.ket} name='ket' onChange={this.handleBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Nilai Sisa</label>
              <Input type='number' value={newBarang.nilaiSisa} name='nilaiSisa' onChange={this.handleBarangInput}/>
            </Card>
            <Card>
              <label>Umur Ekonomis</label>
              <Input type='number' value={newBarang.umurEkonomis} name='umurEkonomis' onChange={this.handleBarangInput}/>
            </Card>
          </SidedCard>
          <Button
            onClick={this.handleSubmit}
            style={{ margin: '8px', width: '100%' }} 
            htmlType='submit' type='primary'>
            Tambah
          </Button>
        </form>
      </div>
      <div style={{ width: '400px', marginLeft: '16px' }}>
        <h2>Daftar Barang Masuk</h2>
        <Table columns={cols} dataSource={barangs} />
      </div>

      <Modal title="Edit Barang Masuk"
        visible={this.state.modVisible}
        onOk={this.handleOk}
        onCancel={this.handleCancel} >

        <form>
          <SidedCard>
            <Card>
              <label>Nama</label> 
              <Select style={{ width: '100%' }} 
                value={editedBarang.nama}
                name='tipeSpek' onChange={this.handleEditNamaBarang}>

                  { this.kodeBarang.map((k, idx) => <Option key={idx} value={idx} >{k.name}</Option>) }
              </Select>
            </Card>
            <Card>
              <label>Kode</label>
              <p style={{ paddingTop: '4px' }}><b>{editedBarang.kode}</b></p>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>No.Reg</label>
              <Input value={editedBarang.reg} name='reg' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Merk</label>
              <Input value={editedBarang.merk} name='merk' onChange={this.handleEditBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Ukuran/CC</label>
              <Input value={editedBarang.ukuran} name='ukuran' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Bahan</label>
              <Input value={editedBarang.bahan} name='bahan' onChange={this.handleEditBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Perolehan</label>
              <Input type='date' value={`${editBarangDate.getFullYear()}-${editBarangDate.getMonth()}-${editBarangDate.getDate()}`} name='tglMasuk' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Nomor</label>
              <div style={{
                display: 'flex',
              }}>
              <Select style={{ width: '100px' }} 
                value={editedBarang.tipeSpek}
                name='tipeSpek' onChange={this.handleEditSpek}>

                  { this.tipe.map((t, idx) => <Option key={idx} value={t} >{t}</Option>) }
              </Select>
              <Input style={{ marginLeft: '8px' }} value={editedBarang.nomorSpek} name='nomorSpek' onChange={this.handleEditBarangInput}/>
              </div>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Cara Perolehan/Status</label>
              <Input value={editedBarang.caraPerolehan} name='caraPerolehan' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Harga</label>
              <Input type='number' value={editedBarang.harga} name='harga' onChange={this.handleEditBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Jumlah</label>
              <Input type='number' value={editedBarang.jml} name='jml' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Keterangan</label>
              <Input value={editedBarang.ket} name='ket' onChange={this.handleEditBarangInput}/>
            </Card>
          </SidedCard>
          <SidedCard>
            <Card>
              <label>Nilai Sisa</label>
              <Input type='number' value={editedBarang.nilaiSisa} name='nilaiSisa' onChange={this.handleEditBarangInput}/>
            </Card>
            <Card>
              <label>Umur Ekonomis</label>
              <Input type='number' value={editedBarang.umurEkonomis} name='umurEkonomis' onChange={this.handleEditBarangInput}/>
            </Card>
          </SidedCard>
          <Button
            onClick={this.handleOk}
            style={{ margin: '8px', width: '100%', display: 'none' }} 
            htmlType='submit' type='primary'>
          </Button>
        </form>
      </Modal>

    </Container>;
  }
}