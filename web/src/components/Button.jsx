import React, { Component } from 'react';

const Button = ({text, dispatch, style, large, fill, action}) => {

    var className = "btn btn-" + style;
    if (large) {
        className += " btn-lg";
    }
    if (fill) {
        className += " btn-block";
    }
   
    return <button type="button" className={className} onClick={() => dispatch(action)}>
        {text}
    </button>;
};

export default Button;
