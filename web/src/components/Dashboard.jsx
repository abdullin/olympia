import React, { Component } from 'react';

import Render from "./Render";



const DashboardCell = ({content, dispatch}) => (
    <div className="col">
      <div className="">
        <Render {...content} dispatch={dispatch} />
        </div>
    </div>
);

const DashboardRow = ({items, dispatch}) => (
    <div className="row">
      {items.map(r => <DashboardCell content={r} dispatch={dispatch}/>)}
    </div>
);


const Dashboard = ({rows = [], dispatch}) => (
    <div className="container">
      {rows.map(r => <DashboardRow items={r} dispatch={dispatch}/>)}
      
        
    </div>
);
    
export default Dashboard;
