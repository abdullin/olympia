import React from 'react';
import { render } from 'react-dom'
import './index.css';
import App from './App';
import ModalDialogSample from './containers/ModalDialogSample';

import 'bootstrap/dist/css/bootstrap.css';
import 'font-awesome/css/font-awesome.css';

import { Provider } from 'react-redux';
import { createStore } from 'redux';
import reducers from './reducers';



const taskFormItems = [
    {type: 'text', label: 'Name', hint: 'new task'},
    {type: 'select', label: 'Priority', options: ['Minor', 'Major']},
];



const taskColumns = [
    {text: "#"},
    {text: "Prio"},
    {text: "Name"},
    {text: "", type:'bool'}
];
const taskRows = [
    {cells:[
        {val:"1"},
        {val:"Normal"},
        {val:"Finish UI prototype"},
        {val:1}
    ]},
    {cells:[
        {val:"2"},
        {val:"Normal"},
        {val:"Prototype View engine server"},
        {val:0}]}

];


var design = {
    type: "window",
    title: "Loading",
    menu:[],
    content:null
};
let store = createStore(
    reducers,
    // preloaded state
    {design:design},
    // plug react dev tools
    window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

var addTaskForm = {
    type: 'form',
    items: [
        {type: 'text', label: 'Name', hint: 'new task'},
        {type: 'select', label: 'Priority', options: ['Minor', 'Major']},
    ]
};

function update1() {
    store.dispatch({
        type: "DESIGN_CHANGED",
        design: { title: "NEW"}
    }); 
}


var ws = new WebSocket("ws://localhost:3001/todo");

ws.onmessage = function(e){
    var msg = JSON.parse(e.data);
    store.dispatch({
        type: "DESIGN_CHANGED",
        design: msg
    });
};

const dispatch = (msg) => {
    console.log("sending", msg)
    var txt = JSON.stringify(msg);
    ws.send(txt);

    
    store.dispatch(msg);
};


render(
        <Provider store={store}>
          <ModalDialogSample dispatch2={dispatch}/>
        </Provider>,
    document.getElementById('root')
);
