import React, { Component } from 'react';

const FormText = ({label, hint, value, onChange, name}) => (
    <div className="form-group">
      <label>{label}</label>
      <input type="text"
             className="form-control"
             placeholder={hint}
             value={value}
             onChange={(e) => onChange(name, e)}/>
    </div>
);

const FormOptions = ({name, options = ['No options provided'], label, value, onChange}) => {

    var rendered = options.map(i => <option>{i}</option>);
    return(
        <div className="form-group">
          <label htmlFor="priority">{label}</label>
          <select className="form-control" value={value} onChange={(e) => onChange(name, e)} >
            <option value="" disabled selected hidden>Please Choose...</option>
            {rendered}
          </select>
        </div>
    );
};

const FormItem = ({type, ...rest}) => {
    switch(type) {
    case 'text':
        return <FormText {...rest} />;
    case 'select':
        return <FormOptions {...rest} />;
    default:
        return (<div className="">Unknown {type}</div>);
    }
};


class Form extends Component{
    constructor(props) {
        super(props);
        this.state = {};
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(name, event) {
        var part = {};
        part[name] = event.target.value;
        this.setState(part);
    }
    
    handleSubmit(event) {
        this.props.onSubmit(this.state);
        event.preventDefault();
    }

    render() {
        return (
            <form>
              {this.props.items.map((i) => <FormItem {...i} onChange={this.handleChange} value={this.state[i.name]} />) }
            </form>

        );
    }
}

export default Form;
