import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { linkTo } from '@storybook/addon-links';

import { Button, Welcome } from '@storybook/react/demo';

import Dialog from '../components/Dialog';

import Scenario from '../components/Scenario';

import Form from '../components/Form';

import Image from '../components/Image';
import Window from '../components/Window';
import DataGrid from '../components/DataGrid';
import Dashboard from '../components/Dashboard';


import XYAxis from '../components/XYAxis';
import BarChart from '../components/BarChart';
import PieChart from '../components/PieChart';
import LineChart from '../components/LineChart';

const taskFormItems = [
    {type: 'text', label: 'Name', hint: 'new task'},
    {type: 'select', label: 'Priority', options: ['Minor', 'Major']},
];


var barChartData = [
                { x: "John", y: 2 },
                { x: "Dom", y: 12 },
                { x: "Tom", y: 5 },
                { x: "Kareem", y: 7 },
                { x: "Tariq", y: 5 },
];

var barChartData2 = [
                { x: "Sweden", y: 122 },
                { x: "Finland", y: 15 },
                { x: "Norway", y: 78 },
    { x: "Poland", y: 5 },
    
                { x: "Russia", y: 20 },
];

var pieChartData  = {
                "React": 2,
                "Relay": 12,
                "GraphQL": 5,
                "Radium": 7,
                "Babel": 5
};


var bubbleChartData = [
                    { name: "Alaa", value: 1 },
                    { name: "Zaid", value: 1 },
                    { name: "Kareem", value: 2 },
                    { name: "Mahmoud", value: 1 },
                    { name: "Tariq", value: 1 },
                    { name: "Shareef", value: 1 },
                    { name: "Tom", value: 41 },
                    { name: "Forest", value: 2 },
                    { name: "John", value: 84 },
                    { name: "Alex", value: 11 },
                    { name: "Donald", value: 7 },
                    { name: "Mark", value: 29 },
                    { name: "Charles", value: 20 },
                    { name: "Quincy", value: 5 },
                    { name: "Alvan", value: 1 },
                    { name: "Don", value: 32 },
                    { name: "Hassan", value: 2 },
                    { name: "Jordan", value: 8 },
                    { name: "Michael", value: 32 },
                    { name: "Steven", value: 5 },
                    { name: "Rafael", value: 2 },
                    { name: "Rick", value: 12 },
];

const taskColumns = [
    {text: "#"},
    {text: "Status"},
    {text: "Name"},
    {text: "", type:'bool'}
];
const taskRows = [
    {cells:[
        {val:"1"},
        {val:"Done"},
        {val:"Finish UI prototype"},
        {val:1}
    ]},
    {cells:[
        {val:"2"},
        {val:"Progress"},
        {val:"Prototype View engine server"},
        {val:0}]},
    
    {cells:[
        {val:"3"},
        {val:"Next"},
        {val:"Clean-up UI component model"},
        {val:0}
    ]},
    
    {cells:[
        {val:"4"},
        {val:""},
        {val:"Code-gen for the component model"},
        {val:0}
    ]},

];


var data = [
    {x: 'K', y: 18},
    {x: 'J', y: 12},
    {x: 'D', y: 16},
    {x: 'S', y: 8},
    {x: 'K2', y: 20},
    {x: 'S2', y: 25},
];


var dashItems = [[
    {
        type: "card",
        header:"Bar Chart",
        footer: "Updated 2 hours ago",
        content: {type:"bar-chart", data:data}
    }, {
        type: "card",
        header: "Pie chart",
        footer: "Updated 3 mins ago",
        content: {type:"pie-chart", width:200, height:200, data:data }
    }, {
        type: "card",
        header: "Another chart",
        content: { type:"line-chart", width:200, height:200, data:data }
    }, {
        type: "card",
        header: "Tasks",
        content: {
            type: "data-grid",
            columns: taskColumns,
            rows: taskRows
        }
    }, {
        type: "card",
        header: "Cats!",
        content: { type: "image", src: "https://http.cat/100"}
    },


    
    
],

                ];


storiesOf('Charts', module)
    .add("Bar chart", () => {


var data = [
  {x: 'K', y: 18},
  {x: 'J', y: 12},
  {x: 'D', y: 16},
  {x: 'S', y: 8},
  {x: 'K2', y: 20},
    {x: 'S2', y: 25},
];

return <XYAxis data={data}
        grid={true}
        xLabel={' '}
        yLabel={' '}
        gridLines={'solid'}>
  <BarChart dataKey='y'/>
            </XYAxis>;
    })
    .add('Pie', () =>{
        var data = [
  {x: 5, y: 63584},
  {x: 10, y: 42839},
  {x: 12, y: 35894},
  {x: 18, y: 58934},
  {x: 25, y: 74323},
  {x: 30, y: 24839},
  {x: 50, y: 12839}
];

var color = [
  "#e1eef6","#ff5f2e","#fcbe32","#004e66","#ff7473","#ffc952","#47b8e0",
  "#34314c","#47b8e0","#47b8e0",
];

//Pie
        return (<PieChart width={350}
          height={300}
          radius={150}
          data={data}
          dataKey='y'
          labelKey='x'
                colors={color}/>);
    })

    .add('Line chart', () => {
var data = [
  {x: 1, y: 6584},
  {x: 2, y: 42839},
  {x: 3, y: 35894},
    {x: 4, y: 12839},
  {x: 5, y: 74323},
  {x: 6, y: 24839},
];


        return (
            <XYAxis data={data}
        grid={true}
        xLabel={' '}
        yLabel={'Bids'}
        gridLines={'solid'}>

  <LineChart dataKey='y'/>
            </XYAxis>);

    });


var scenario = {
    title: "Task management system with 4 tasks and some progress changes",
    blocks : [
        {
            title: "GIVEN",
            cards: [
                {
                    title: "Tenant configured with task management",
                    code:
`{ event: 'tenant/created', tenant: 1, name: 'any'}
{ event: 'feature/added', tenant: 1, type: 'task-management', config: {…} },
{ event: 'user/created', tenant: 1, user: 1, capabilities: [ 'task/admin' ], name: 'New User' }
`
                },
                {
                    title: "Tasks added to the system",
                    code:
`{ event: 'task/added', tenant: 1, user: 1, task: 1, title: 'Finish UI prototype'}
{ event: 'task/added', tenant: 1, user: 1, task: 2, title: 'Prototype View engine server'}
{ event: 'task/added', tenant: 1, user: 1, task: 3, title: 'Clean-up UI component model'}
{ event: 'task/added', tenant: 1, user: 1, task: 4, title: 'Code-gen for the component model'}`
                },

                {
                    title: "Tasks progress changed by the user",
                    code:
`{ event: 'task/completed', tenant: 1, user: 1, task: 1}
{ event: 'task/started', tenant: 1, user: 1, task: 2}
{ event: 'task/scheduled', tenant: 1, user: 1, task: 3}`
                }

                
            ]
        },
        {
            title: "WHEN",
            cards: [
                {
                    title: "User 1/1 starts a new session",
                }
            ],
        },
        {
            title: "THEN",
            cards: [
                
                {
                    title: "Home screen should display grid with 4 tasks",
                    render: {
                        type: "window",
                        title: "{any}",
                        nav: [
                        ],
                        content:{type:'data-grid', columns: taskColumns,rows:taskRows}
                        
                    }
                },
                {
                    title: "Home screen should have nav 'Tasks (3)'",
                    render: {
                        type: "window",
                        title: "{any}",
                        nav: [
                            {text: "Tasks (3)", active:1}
                        ],
                        content:{type:'placeholder', text: ''}
                        
                    },
                    text: "NOTE: we display number of incomplete AND active tasks in the nav"
                },

                
            ]
        }
        

    ]



};



storiesOf('Scenario', module)
    .add('Sample', () => <Scenario {...scenario} />);

storiesOf('Dashboard', module)
    .add('Simple', () => <Dashboard rows={dashItems}/>);


storiesOf('DataGrid', module)
    .add('Main', () => <DataGrid columns={taskColumns} rows={taskRows}/>);

storiesOf('Window', module)
    .add('Main', () => {
        var props = {
            title: "Main window",
            menu:[{title:"Admin", active: 1},{title:"Help"}],
            content:{type:"data-grid", columns:taskColumns, rows: taskRows, canEdit:1, canDelete:1},
            nav:[
                {text:"Tasks (1)", active: 1},
                {text:"Dashboards"},
                {text:"Reports"},
                {text:"Audit"}
            ]
            
        };
        return <Window {...props} />;});




storiesOf('Dialog', module)
    .add('Basic story', () => {

        var options = {
            title: "Create Task",
            ok: "Save",
            cancel: "Cancel",
            content: {
                type: 'form',
                items: taskFormItems
            }
        };
        
        return <Dialog {...options}/>;
    });

storiesOf('Image', module)
    .add("chart", () => {
        return <Image src="https://github.com/vdobler/chart/blob/master/example/bestof.png?raw=true" />;

    });





storiesOf('Form', module)
    .add("Sample form", () => {

        var items = [
            {type:'text', label:'Name'}

        ];

        var props = {
            items: [
                {type: 'text', label: 'Name', hint: 'new task', name:'name'},
                {type: 'select', label: 'Priority', options: ['Minor', 'Major'], name:'priority'},
            ]

        };


        return <Form {...props}/>;});

