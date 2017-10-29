const initialState = { title: 'hello!', content: {}}


const design = (state = initialState, action) => {
  switch (action.type) {
    case 'DESIGN_CHANGED':
      return action.design;
    default:
      return state
  }
}

export default design;
