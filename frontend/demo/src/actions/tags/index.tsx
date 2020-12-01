import {createAction} from 'typesafe-actions';
import _, {Dispatch} from 'redux';

import {doFetchFromBackend} from '../../api/airyConfig';
import {Tag, TagPayload, CreateTagRequestPayload} from '../../model/Tag';

export const UPSERT_TAG = 'UPSERT_TAG';
export const DELETE_TAG = 'DELETE_TAG';
export const EDIT_TAG = 'EDIT_TAG';
export const ERROR_TAG = 'ERROR_TAG';
export const ADD_TAGS_TO_STORE = 'ADD_TAGS_TO_STORE';
export const SET_TAG_FILTER = 'SET_TAG_FILTER';

export function fetchTags(tags: Tag[]) {
  return {
    type: ADD_TAGS_TO_STORE,
    tagData: tags,
  };
}

export function addTag(tag: Tag) {
  return {
    type: UPSERT_TAG,
    tagData: tag,
  };
}

export function editedTag(id: string, name: string, color: string, count: number) {
  return {
    type: EDIT_TAG,
    tagData: {
      id,
      name,
      color,
      count,
    },
  };
}

export function errorTag({status}) {
  return {
    type: ERROR_TAG,
    tagData: {
      status,
    },
  };
}

export function deleteConversationTag(tagId: string) {
  return {
    type: DELETE_TAG,
    tagData: {
      tag_id: tagId,
    },
  };
}

export function getTags(query: string = '') {
  return function(dispatch: Dispatch<any>) {
    return doFetchFromBackend('tags.list').then((response: Tag[]) => {
      console.log(response)
      dispatch(fetchTags(response));
    });
  };
}

export function createTag(requestPayload: CreateTagRequestPayload) {
  return async (dispatch: Dispatch<any>) => {
    return doFetchFromBackend('tags.create', requestPayload)
      .then((response: TagPayload) => {
        const tag: Tag = {
          id: response.id,
          name: requestPayload.name,
          color: requestPayload.color,
          count: 0,
        };
        dispatch(addTag(tag));
        return true;
      })
      .catch(error => {
        dispatch(errorTag(error));
        return false;
      });
  };
}

export function updateTag(tagId: string, name: string, color: string, count: number) {
  return function(dispatch: Dispatch<any>) {
    doFetchFromBackend('tags.update', {
      id: tagId,
      name: name,
      color: color,
    }).then(tag => dispatch(editedTag(tagId, name, color, count)));
  };
}

export function deleteTag(tagId: string) {
  return function(dispatch: Dispatch<any>) {
    doFetchFromBackend('tags.delete', {
      id: tagId,
    }).then(() => {
      dispatch(deleteConversationTag(tagId));
    });
  };
}

export function filterTags(filter) {
  return function(dispatch: Dispatch<any>) {
    dispatch({
      type: SET_TAG_FILTER,
      payload: filter,
    });
  };
}
