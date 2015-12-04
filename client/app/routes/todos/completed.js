import TodosIndex from './index';

export default TodosIndex.extend({
  model() {
    return this.store.filter('todo', function(todo) {
      return todo.get('isCompleted');
    });
  }
})
