import React, { Component } from 'react';
import Render from "./Render";

const Header = ({text}) => {
    if (text != null) {
        return <div className="card-header">{text}</div>;
    }
    return null;
};

const Footer = ({text}) => {
    if (text == null)
        return null;

    return (
        <p className="card-text">
          <small className="text-muted">
            {text}
          </small>
        </p>); 
};

const Card = ({content, dispatch, title, header, footer}) => (
    <div className="card mb-3">
      <Header text={header}/>
      <div className="card-body">
        <Render {...content} dispatch={dispatch} />
        <Footer text={footer} />
      </div>
    </div>
);

export default Card;
