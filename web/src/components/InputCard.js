import React from 'react';
import styled from 'styled-components';
import { Input } from 'antd';

const Card = styled.div`
  margin-bottom: 8px;
  margin-right: 8px;
  width: ${({ width }) => width ? width+'px' : '300px'};
`

export const InputCard = ({ width, type = 'text', label = '', value, name = '', onChange }) => {
  return <Card width={width}>
  <label>{label}</label>
  <Input type={type} value={value} name={name} onChange={(e) => onChange(e)}/>
  </Card>
}