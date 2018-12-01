import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { hostname: 'n/a' };
  }

  componentDidMount() {
    fetch('api/hostname')
    .then(response => response.text())
    .then(hostname => {
      this.setState({ hostname: hostname });
    });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            {this.state.hostname}
          </p>
        </header>
      </div>
    );
  }
}

export default App;
