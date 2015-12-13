import DS from 'ember-data';

export default DS.JSONAPISerializer.extend({
  normalizeFindAllResponse(store, primaryModelClass, payload, id, requestType) {
    return {
      data: payload.map((item) => {
        let newItem = this.normalize(primaryModelClass, item).data
        return newItem;
      })
    }
  },
  extractId(modelClass, resourceHash) {
    return resourceHash.id
  },
  extractAttributes (modelClass, resourceHash) {
    delete  resourceHash.id;
    delete resourceHash.type;
    return resourceHash;
  }
});
