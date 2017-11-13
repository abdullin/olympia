import React, { Component } from 'react';


const Image = ({src, dispatch, action}) => {

    if (action == null) {
        return <img className="img-fluid" src={src} />;
    }
    
    return <img className="img-fluid" src={src} onClick={() => dispatch(action) }/>;
};

export default Image;
