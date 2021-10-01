import React, {Component} from 'react';
import { StatusBar } from 'expo-status-bar';
import { Platform, NativeModules } from 'react-native';
import Main from './src/Main';

export default class App extends Component {
  render() {
    return (
      <>
        <Main />
        <StatusBar style="auto"/>
      </>
    );
  }
  
}

