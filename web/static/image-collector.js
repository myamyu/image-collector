class ImageCollectorElement extends HTMLElement {
  constructor() {
    super();

    const shadowRoot = this.attachShadow({mode: 'open'});
    shadowRoot.innerHTML = `
      <style>
        .result-images {
          display: flex;
          flex-wrap: wrap;
          width: 98vw;
          margin: 0;
          padding: 0;
          justify-content: space-between;
        }
        .result-img {
          margin: 0 0 1vw;
          padding: 0;
          display: block;
          width: 19vw;
          height: 19vw;
        }
        .page-link {
          display: block;
          width: 100%;
          height: 100%;
          text-decoration: none;
        }
        .result-img img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
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
      <div class="results-num"><span class="results-num-val">${images.length}</span>件</div>
      <ul class="result-images"></ul>
    `;
    const resultImages = this._results.querySelector('.result-images');
    images.forEach((img) => {
      const li = document.createElement('li');
      li.classList.add('result-img')
      li.innerHTML = `
        <a href="${img.web_page_url}" class="page-link"><img src="${img.thumb_url}" alt="${img.text}"></a>
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

export {ImageCollectorElement}
