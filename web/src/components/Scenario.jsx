import React, { Component } from 'react';

import Render from './Render';

const Card = ({title, text, code, render}) => {
    var titleEl, codeEl, textEl, renderEl;
    
    if (title != null) {
        titleEl = <h6 className="card-header">{title}</h6>;
    }
    if (text != null) {
        textEl = (
            <div className="card-body">
              <p className="card-text">{text}</p>
            </div>
        );
    }
    if (code != null) {
        codeEl = (
            <div className="card-body">
              <pre className="mb-0"><code>{code}</code></pre>
            </div>
        );
    }
    if (render != null) {
        renderEl = (<Render {...render} />);
    }
    
    return (
        <div className="card mb-3">
          {titleEl}
          {textEl}
          {codeEl}
          {renderEl}
        </div>);
};

const Block = ({title, cards = []}) => (
    <div className="alert alert-primary">
      <h4>{title}</h4>
      {cards.map(Card)}
    </div>
);

const Scenario = ({title, blocks = []}) => (
    <div className="m-3">
      <h1>{title}</h1>
      {blocks.map(Block)}
    </div>
);

export default Scenario;
