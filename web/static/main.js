import {ImageCollectorElement} from './image-collector.js';

customElements.define('image-collector', ImageCollectorElement);

((d) => {
  d.getElementById('searchButton').addEventListener('click', (e) => {
    const result = d.getElementById('searchResults');
    result.setAttribute('site-name', d.querySelector('[name=s]').value);
    result.setAttribute('search-word', d.querySelector('[name=q]').value);
    result.setAttribute('limit', '100');

    result.search();
  });
})(document);
