import React, { Component } from 'react';
import Render from "./Render";

const Row = ({cols, dispatch}) => {
    return <div className="row">
        {cols.map(c => <Col {...c} dispatch={dispatch}/>)}
    </div>;
};

const Col = ({content, dispatch, steps}) => {

    var style = "col";
    if (steps) {
        style += " col-" + steps;
    }
    return <div className={style}>
        <Render {...content} dispatch={dispatch} />
        </div>;
};

const Grid = ({rows, dispatch}) => (
    <div className="container">
      {rows.map(r => <Row {...r} dispatch={dispatch} />)}
    </div>
);

export default Grid;
