import { connect } from 'react-redux';

import React from 'react';
import Window from "../components/Window";
import Dialog from "../components/Dialog";

const mapStateToProps = state => {
    return state.design;
};


const RootContainer = ({type, dispatch2, ...props}) => {
    switch (type) {
    case "window":
        return <Window {...props} dispatch={dispatch2}/>;
    case "dialog":
        return <Dialog {...props} dispatch={dispatch2} />;
    default:
        return <h1>Unknown type {type}</h1>;
    }
    
};

const ModalDialogSample = connect(
  mapStateToProps
)(RootContainer);


export default ModalDialogSample;
