import { configure } from '@storybook/react';
import 'bootstrap/dist/css/bootstrap.css';
import 'font-awesome/css/font-awesome.css';
function loadStories() {
  require('../src/stories');
}

configure(loadStories, module);
