import React, { Component } from 'react';


import XYAxis from './XYAxis';
import BarChart from './BarChart';
import PieChart from './PieChart';
import LineChart from './LineChart';
import Responsive from './Responsive';

import Form from './Form';
import DataGrid from './DataGrid';
import Card from './Card';
import Window from './Window';

import Image from './Image';

import Button from './Button';

import Grid from './Grid';

import Placeholder from './Placeholder';

// Render component takes any json design and returns full model
const Render = ({type, responsive, width, height, ...rest}) => {

    if (type == null) {
        return null;
    }


    
    switch(type){
    case 'image':
        return <Image {...rest} />;
        
    case 'bar-chart':

        var chart = (
            <XYAxis
              {...rest}
              grid={true}
              xLabel={' '}
              yLabel={' '}
              gridLines={'solid'}
              height={height}
              width={width}
              >
              
  <BarChart dataKey='y'/>
            </XYAxis> 
        );
        return chart;
        
    case 'line-chart':
        return (

            <XYAxis {...rest}
                    grid={true}
                    xLabel={' '}
                    yLabel={' '}
                    gridLines={'solid'}

                    height={height}
                    width={width}
                    >

  <LineChart dataKey='y'/>
</XYAxis>

        );

    case 'pie-chart':
        
var color = [
  "#e1eef6","#ff5f2e","#fcbe32","#004e66","#ff7473","#ffc952","#47b8e0",
  "#34314c","#47b8e0","#47b8e0",
];
        
        return <PieChart
        dataKey='y'
        labelKey='x'
        radius={150}
        colors={color}
        
              height={height}
              width={width}
        {...rest}/>;
        
    case 'form':
        return (<Form {...rest} />);
    case 'data-grid':
        return (<DataGrid {...rest}/>);
    case 'card':
        return <Card {...rest} />;
    case 'window':
        return <Window {...rest} />;
    case 'button':
        return <Button {...rest} />;
    case 'placeholder':
        return <Placeholder {...rest} />;
    case 'grid':
        return <Grid {...rest} />;
        
    default:
        return (<p>Unknown render type {type}</p>);
    }
};

export default Render;
