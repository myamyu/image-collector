class ImageCollectorElement extends HTMLElement {
  constructor() {
    super();

    const shadowRoot = this.attachShadow({mode: 'open'});
    shadowRoot.innerHTML = `
      <style>
      </style>
      <div class="search-result"></div>
    `;
    this._results = shadowRoot.querySelector('.search-result');

    this.search();
  }

  waiting() {
    this._results.innerHTML = `
      <div class="waiting">
        <p>処理中です...お待ちください</p>
      </div>
    `;
  }

  result(images) {
    this._results.innerHTML = `
      <h3><span class="site-name">${this._siteName}</span>で<span class="search-word">${this._searchWord}</span>の検索結果</h3>
      <hr>
      <div class="results-num"><span class="results-num-val">${images.length}</span>件</div>
      <ul class="result-images"></ul>
    `;
    const resultImages = this._results.querySelector('.result-images');
    images.forEach((img) => {
      const li = document.createElement('li');
      li.classList.add('result-img')
      li.innerHTML = `
        <div><img src="${img.thumb_url}" alt="${img.text}"></div>
      `;
      resultImages.appendChild(li);
    });
  }

  search() {
    this._siteName = this.getAttribute('site-name') || 'twitter.com';
    this._searchWord = this.getAttribute('search-word') || 'スタバなう';
    this._limit = 0|this.getAttribute('limit') || 30;
    this.waiting();

    const url = `/image-collector?s=${encodeURIComponent(this._siteName)}&q=${encodeURIComponent(this._searchWord)}&l=${this._limit}`;
    fetch(url)
      .then((res) => {
        return res.json();
      })
      .then((res) => {
        this.result(res);
      })
  }
}

customElements.define('image-collector', ImageCollectorElement);

((d) => {
  d.getElementById('searchButton').addEventListener('click', (e) => {
    const result = d.getElementById('searchResults');
    result.setAttribute('site-name', d.querySelector('[name=s]').value);
    result.setAttribute('search-word', d.querySelector('[name=q]').value);
    result.setAttribute('limit', d.querySelector('[name=l]').value);

    result.search();
  });
})(document);
