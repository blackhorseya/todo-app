function handleResponse(resp) {
  return resp.text().then((body) => {
    const data = body && JSON.parse(body);
    if (resp.status !== 200) {
      const error = (data && data.message) || resp.statusText;
      return Promise.reject(error);
    }

    return data;
  });
}

function list() {
  const reqOpt = {
    method: 'GET',
  };

  return fetch(`${process.env.REACT_APP_API_URL}api/v1/tasks?size=10`, reqOpt)
      .then(handleResponse);
}

export const taskServices = {
  list,
};
