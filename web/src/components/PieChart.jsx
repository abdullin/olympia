import * as d3 from "d3";
import React from 'react';
import ReactFauxDOM from 'react-faux-dom';
import PropTypes from 'prop-types';

class PieChart extends React.Component {
  render () {
    const data = this.props.data,
          dataKey = this.props.dataKey || 'y',
          labelKey = this.props.labelKey || 'x',
          width= this.props.width || 350,
          height = this.props.height || 300,
          radius = Math.min(width, height) /2.1,
          donut = this.props.donut,
          colorFill = (donut) ? radius - radius / donut : 0,
          textColor = this.props.textColor || '#fff',
          colors = this.props.colors || ["#e1eef6","#ff5f2e","#fcbe32","#004e66",
          "#ff7473","#ffc952","#47b8e0","#34314c","#47b8e0","#47b8e0"],
          color = d3.scaleOrdinal()
            .range(colors);

    var arc = d3.arc()
        .outerRadius(radius - 10)
        .innerRadius(colorFill);

    var labelArc = d3.arc()
        .outerRadius(radius - 40)
        .innerRadius(radius - 40);

    var pie = d3.pie()
        .sort(null)
        .value((d) => d[dataKey]);

    var chart = d3.select(ReactFauxDOM.createElement('svg'))
        .attr('class', 'PieChart')
        .attr('width', width)
        .attr('height', height)
        .style('display', 'inline-block');

    var g = chart.append('g')
        .attr('transform', `translate(${width / 2}, ${height / 2})`)
        .style('display', 'inline-block');

    var slice = g.selectAll('.arc')
        .data(pie(data))
        .enter().append('g')
        .attr('class', 'arc');

    slice.append('path')
        .attr('d', arc)
        .style('fill', (d) => color(d.data[labelKey]));

    slice.append('text')
        .attr('transform', (d) => `translate(${labelArc.centroid(d)})`)
        .attr('dy', '.35em')
        .text((d) => d.data[labelKey])
        .style('font', '11px sans-serif')
        .style('fill', textColor)
        .style('text-anchor', 'middle');

      return chart.node().toReact();
  }
}


PieChart.propTypes = {
  width: PropTypes.number,
  height: PropTypes.number,
  colors: PropTypes.array,
  donut: PropTypes.number,
  data: PropTypes.array,
  dataKey: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
  labelKey: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
  textColor: PropTypes.string
};

export default PieChart;
