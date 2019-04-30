import React from 'react';
import { Input, Select, Button } from 'antd';
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

export class BarangMasuk extends React.Component {
  _isMounted = false
  tipe = ['Polisi', 'Mesin']
  state = {
    barangs: [],
    newBarang: {
      kode: '', 
      nama: '', 
      reg: '',
      merk: '',
      ukuran: '',
      bahan: '',
      tglMasuk: 0,
      tipeSpek: this.tipe[0],
      nomorSpek: '',
      caraPerolehan: '',
      jml: 0,
      harga: 0,
      ket: '',
    }
  }

  componentDidMount() {
    this._isMounted = true
  }

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
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    if (!this._isMounted) return;
    this.setState({
      newBarang: {
        ...this.state.newBarang,
        [name]: value,
      }
    });
    console.log(value);
  }

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
            <Input value={newBarang.tglMasuk} name='tglMasuk' onChange={this.handleBarangInput}/>
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
            <Input value={newBarang.harga} name='harga' onChange={this.handleBarangInput}/>
          </Card>
        </SidedCard>
        <SidedCard>
          <Card>
            <label>Keterangan</label>
            <Input value={newBarang.ket} name='ket' onChange={this.handleBarangInput}/>
          </Card>
        </SidedCard>
        <Button style={{ margin: '8px', width: '100%' }} htmlType='submit' type='primary'>Tambah</Button>
      </form>
    </Container>;
  }
}