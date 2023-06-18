import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalidades.css';

export default class Personalidades extends Component {
    state = {
        personalidades: []
    }

    componentDidMount() {
        axios.get('http://localhost:3333/personalities')
            .then(res => {
                const personalidades = res.data;
                this.setState({ personalidades })
            })
            .catch(error => {
              console.error(error)
            })
    }

    render() {
        return (
            <div>
                {this.state.personalidades.map((person, id )=>
                    <div className="CardPersonalidades" key={id}>
                        <h3>{person.name}</h3>
                        <p>{person.history}</p>
                    </div>)}
            </div>
        );
    }
}