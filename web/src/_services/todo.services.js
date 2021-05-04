const endpoint = `${process.env.REACT_APP_API_ENDPOINT || ''}`;

export const todoService = {
  list,
  add,
  remove,
};

function list(start, end) {
  const opts = {
    method: 'GET',
  };

  return fetch(`${endpoint}/api/v1/tasks?start=${start}&end=${end}`, opts).then(handlerResp);
}

function add(task) {
  const opts = {
    method: 'POST',
    body: JSON.stringify(task),
  };

  return fetch(`${endpoint}/api/v1/tasks`, opts).then(handlerResp);
}

function remove(id) {
  const opts = {
    method: 'DELETE',
  };

  return fetch(`${endpoint}/api/v1/tasks/${id}`, opts).then((response) => {
    return response.text().then(text => {
      const resp = text && JSON.parse(text);

      if (response.statusCode >= 400) {
        const error = (resp && resp.msg) || response.statusText;
        return Promise.reject(error);
      }

      return id;
    });
  });
}

function handlerResp(response) {
  return response.text().then(text => {
    const resp = text && JSON.parse(text);

    if (response.statusCode >= 400) {
      const error = (resp && resp.msg) || response.statusText;
      return Promise.reject(error);
    }

    return resp.data;
  });
}