import React, { Component } from 'react';
import Render from "./Render";



const NavItem = ({title, active, action, dispatch}) => (
          <li className={active?"nav-item active":"nav-item"}>
            <a className="nav-link" href="#"
               onClick={e => {e.preventDefault(); dispatch(action);}}> {title}
            </a>
          </li>
);

const SplitContainer = ({left, right}) => (
    <div className="container-fluid">
      <div className="row">
        <div className="col-sm-4 col-md-3 ">
          {left}
        </div>
        <div className="col-sm-8 col-md-9">
          {right}
        </div>
      </div>
    </div>


);

const NavBar = ({items = [], title, dispatch}) => (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      <div className="collapse navbar-collapse" id="navbarCollapse">
        <ul className="navbar-nav mr-auto">
          {items.map(i => <NavItem {...i} dispatch={dispatch}/>)}
        </ul>
      </div>
      
      <a className="navbar-brand" href="#">{title}</a>
      
    </nav>
);


const NavTreeItem = ({text, state, dispatch, action}) => (

     
    <button type="button"
            disabled={state=="disabled"}
            onClick={e => {e.preventDefault(); dispatch(action);}}
            className={"list-group-item list-group-item-action " + state}>{text}</button>
);
const NavTree = ({items = [], dispatch}) => (
    <ul className="list-group mb-3">
      {items.map(i => <NavTreeItem {...i} dispatch={dispatch} />)}
    </ul>
);


const Window = ({content, title, menu, nav, dispatch}) => (
    <div className="">
      <NavBar items={menu} title={title} dispatch={dispatch}/>
      <SplitContainer right={<Render {...content} dispatch={dispatch}/>}
                      left={<NavTree items={nav} dispatch={dispatch} />} />
    </div>
);

export default Window;
