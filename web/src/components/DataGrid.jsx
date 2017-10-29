import React, { Component } from 'react';



const HeadColumn = ({text}) => (<th>{text}</th>);


const Cell = ({val, column}) => {

    if (column == null) {
        return <td className="table-danger">{val}</td>;
    }

    switch (column.type) {
    case 'bool':
        return val ?
            <td><i className="fa fa-check"/></td> :
            <td/>;
    
        
    default: return <td>{val}</td>;
    }


};


const RowEditor = ({canEdit, canDelete}) => {
    var opts = [];
    if (canEdit) {
        opts.push(<i className="btn fa fa-pencil"/>);
    }
    if (canDelete) {
        opts.push(<i className="btn fa fa-trash"/>);
    }
    return opts.length == 0 ? null : <td>{opts}</td>;
};

const Row = ({cells = [], columns, ...props}) => (
    <tr>
      {cells.map((cell, i) => <Cell {...cell} column={columns[i]}/>)}
        <RowEditor {...props}/>
    

    </tr>
);

const task = {
    type: "show-add-task-form",
    args: {name: "new task", priority: "Major"}
};


const Toolbar = ({dispatch}) => (


    
    <div className="btn-toolbar mb-3">
      <div className="btn-group mr-2 btn-group-sm">
        <button type="button" className="btn btn-light" onClick={() => dispatch(task)}><i className='fa fa-plus'/> Add</button>
        
        
      </div>


      
      <div className="btn-group mr-2 btn-group-sm float-right">
      
        <button type="button" className="btn btn-light"><i className='fa fa-download'/> Download</button>
      </div>
      
      <div className="btn-group mr-2 btn-group-sm">
        <button type="button" className="btn-light btn btn-cog dropdown-toggle"><i className='fa fa-cog'/> Manage View</button>
        
      </div>
    </div>
);

const DataGrid = ({columns = [], rows = [], canEdit, canDelete, dispatch}) => (
    <div className="temporary">
    <table className="table table-sm">
      <thead>
        <tr>
          {columns.map(HeadColumn)}
          {(canEdit || canDelete) ? <th/> : null }
        </tr>
      </thead>
      <tbody>
        {rows.map((row, idx) => <Row columns={columns} {...row} canEdit={canEdit} canDelete={canDelete}/>)}
      </tbody>
        </table>
        </div>
);


export default DataGrid;
