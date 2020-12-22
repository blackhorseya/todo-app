const apiUrl = process.env.REACT_APP_API_URL || '';

function handleResponse(resp) {
  return resp.text().then((body) => {
    const data = body && JSON.parse(body);
    if (400 <= resp.status) {
      const error = (data && data.message) || resp.statusText;
      return Promise.reject(error);
    }

    return data;
  });
}

function list(page, size) {
  const reqOpt = {
    method: 'GET',
  };

  return fetch(`${apiUrl}api/v1/tasks?page=${page}&size=${size}`,
      reqOpt).then(handleResponse);
}

function add(title) {
  const reqOpt = {
    method: 'POST',
    body: JSON.stringify({title}),
  };

  return fetch(`${apiUrl}api/v1/tasks`, reqOpt).then(handleResponse);
}

function remove(id) {
  const reqOpt = {
    method: 'DELETE',
  };

  return fetch(`${apiUrl}api/v1/tasks/${id}`, reqOpt).then(handleResponse);
}

function changeStatus(id, completed) {
  const reqOpt = {
    method: 'PATCH',
  };

  return fetch(`${apiUrl}api/v1/tasks/${id}?completed=${completed}`, reqOpt)
      .then(handleResponse);
}

export const taskServices = {
  list,
  add,
  remove,
  changeStatus,
};
