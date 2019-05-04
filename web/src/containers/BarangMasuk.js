import React from 'react';
import axios from 'axios';
import { Input, Select, Button, notification } from 'antd';
import styled from 'styled-components';

const { Option } = Select;

const Container = styled.div`
  padding: 16px;
  box-sizing: border-box;
`

const Card = styled.div`
  padding: 8px;
  width: 300px;
`

const SidedCard = styled.div`
  display: flex;
`

const tipe = ['Polisi', 'Mesin']
const initNewBarang = {
  kode: '', 
  nama: '', 
  reg: '',
  merk: '',
  ukuran: '',
  bahan: '',
  tglMasuk: null,
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
  bebanPenyusutan: 0.0
}

export class BarangMasuk extends React.Component {
  _isMounted = false
  tipe = ['Polisi', 'Mesin']
  state = {
    barangs: [],
    newBarang: { ...initNewBarang },
  }

  componentDidMount() {
    this._isMounted = true
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

  asyncWrapper = (fn) => {
    fn()
      .catch(e => this.openNotificationWithIcon('error', e.message))
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

    // const keys = Object.keys(newBarang);
    // if (!keys.every(k => newBarang[k])) {
    //   this.openNotificationWithIcon('error', 'Field cannot be empty or 0');
    //   return;
    // }

    try {
      const { data } = await axios('/api/barang', {
        method: 'POST',
        data: {
          ...newBarang,
          tglMasuk: (new Date(newBarang.tglMasuk)).getTime(),
        }
      });
      this.setState({ newBarang: initNewBarang })
      console.log(data);
    } catch (e) {
      this.openNotificationWithIcon('error', 'Field cannot be empty');
    }
  };

  render() {
    const { newBarang } = this.state;

    return <Container>
      <h2>Barang Masuk</h2>
      <form>
        <SidedCard>
          <Card>
            <label>Nama</label>
            <Input value={newBarang.nama} name='nama' onChange={this.handleBarangInput}/>
          </Card>
          <Card>
            <label>Kode</label>
            <Input value={newBarang.kode} name='kode' onChange={this.handleBarangInput}/>
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
        <SidedCard>
          <Card>
            <label>Umur Penggunaan</label>
            <Input type='number' value={newBarang.umurPenggunaan} name='umurPenggunaan' onChange={this.handleBarangInput}/>
          </Card>
          <Card>
            <label>Nilai Buku</label>
            <Input type='number' value={newBarang.nilaiBuku} name='nilaiBuku' onChange={this.handleBarangInput}/>
          </Card>
        </SidedCard>
        <SidedCard>
          <Card>
            <label>Beban Penyusutan</label>
            <Input type='number' value={newBarang.bebanPenyusutan} name='bebanPenyusutan' onChange={this.handleBarangInput}/>
          </Card>
        </SidedCard>
        <Button
          onClick={this.handleSubmit}
          style={{ margin: '8px', width: '100%' }} 
          htmlType='submit' type='primary'>
          Tambah
        </Button>
      </form>
    </Container>;
  }
}