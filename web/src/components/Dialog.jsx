import React, { Component } from 'react';

import Render from './Render';




const Button = ({text, action, style, dispatch}) => (
    <a className={"btn "+style}
       href="#"
       onClick={e => {e.preventDefault(); dispatch(action);}}
      >{text}</a>
);

const DialogFooter =({acceptButton, cancelButton, dispatch}) => {
    var f = [];
    
    if (cancelButton != null) {
        f.push(<Button dispatch={dispatch} {...cancelButton}/>);
    }
    if (acceptButton != null) {
        f.push(<Button
               dispatch={dispatch}
               style="btn-primary"
               {...acceptButton}/>);
    }

    if (f.length==0) {
        return null;
    }
    return (<div className="float-right">{f}</div>);
    
};

const Dialog = ({title, content, dispatch, ...props}) => (

          <div className="card rounder-0">
            <h4 className="card-header">{title}</h4>
            <div className="card-body">
              <Render {...content} dispatch={dispatch}/>
              <DialogFooter {...props} dispatch={dispatch}/>
            </div>
          </div>

);


export default Dialog;
