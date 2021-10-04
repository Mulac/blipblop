import React, {Component} from 'react';
import { StatusBar } from 'expo-status-bar';
import { LogBox } from 'react-native';
import Main from './src/Main';

LogBox.ignoreLogs(['Reanimated 2']);

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

