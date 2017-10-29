import React from 'react';
import * as d3 from 'd3';
import ReactFauxDOM from 'react-faux-dom';

import PropTypes from 'prop-types';
class LineChart extends React.Component {
  render () {
    const data = this.props.data,
          xDataKey = this.props.xDataKey,
          yDataKey = this.props.yDataKey,
          dataKey = this.props.dataKey || yDataKey,
          width= this.props.width || 350,
          height = this.props.height || 300,
          innerW = width - 60,
          innerH = height - 60,
          xScale = this.props.xScale,
          yScale = this.props.yScale,
          color = this.props.color || '#47b8e0',
          pointBorderColor = this.props.pointBorderColor || '#ffc952',
          pointColor = this.props.pointColor || '#fff';


    var line = d3.line()
        .curve(d3.curveMonotoneX)
        .x((d) => xScale(d[xDataKey]))
        .y((d) => yScale(d[dataKey]));

      var planeElement = ReactFauxDOM.createElement('svg');
    var plane = d3.select(planeElement)
        .attr('class', 'LineChart')
        .attr('width', width)
        .attr('height', height)
        .style('z-index', this.props.zIndex)
        .style('position', 'absolute');

    var g = plane.append('g')
        .attr('class', 'plane')
        .attr('width', innerW)
        .attr('height', innerH)
        .attr('transform', `translate(50, 20)`)
        .style('display', 'inline-block');

    g.append('path')
        .datum(data)
        .attr('class', 'line')
        .attr('d', line)
        .style('fill', 'none')
        .style('stroke', color)
        .style('stroke-width', 2);

    g.selectAll('circle')
      .data(data)
      .enter().append('circle')
      .attr('class', 'circle')
      .attr('cx', (d) => xScale(d[xDataKey]))
      .attr('cy', (d) => yScale(d[dataKey]))
      .attr('r', 3)
      .style('fill', pointColor)
      .style('stroke', pointBorderColor);

    return plane.node().toReact();
  }
}

LineChart.propTypes = {
  width: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  height: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  data: PropTypes.array,
  dataPoints: PropTypes.bool,
  dataKey: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
  color: PropTypes.string,
  pointColor: PropTypes.string,
  pointBorderColor: PropTypes.string
}

export default LineChart;
