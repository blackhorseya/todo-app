const endpoint = `${process.env.REACT_APP_API_ENDPOINT || ''}`;

export const todoService = {
  list,
  add,
  remove,
  changeStatus,
};

function list(start, end) {
  const opts = {
    method: 'GET',
  };

  return fetch(`${endpoint}/api/v1/tasks?start=${start}&end=${end}`, opts).then((resp) => {
        return resp.json().then(body => {
          if (!resp.ok) {
            const error = (body && body.msg) || resp.statusText;
            return Promise.reject(error);
          }

          return {
            data: body.data,
            total: resp.headers.get('X-Total-Count'),
          };
        });
      });
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

function changeStatus(id, status) {
  const opts = {
    method: 'PATCH',
    body: JSON.stringify({completed: status}),
  };

  return fetch(`${endpoint}/api/v1/tasks/${id}/status`, opts).then(handlerResp);
}

function handlerResp(response) {
  console.log(response.headers.get('X-Total-Count'));

  return response.json().then(body => {
    if (!response.ok) {
      const error = (body && body.msg) || response.statusText;
      return Promise.reject(error);
    }

    return body.data;
  });
}