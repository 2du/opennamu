"use strict";

let ringo_save_data = '';
let ringo_open = 0;
let ringo_menu_list = [
    'recent_cel',
    'other_cel',
    'user_cel',
    'add_cel'
];

function ringo_opening(data) {
    let element = [data];
    
    for(let for_a in ringo_menu_list) {
        if(ringo_menu_list[for_a] + '_in' !== data) {
            element.push(ringo_menu_list[for_a] + '_in');
        }
    }

    if((document.getElementById(element[0]).style.display === 'none' && ringo_open === 0) || ringo_save_data !== data) {
        document.getElementById(element[0]).style.display = 'block';

        for(let for_a in element) {
            if(for_a !== '0') { 
                if(document.getElementById(element[for_a]) !== null) {
                    document.getElementById(element[for_a]).style.display = 'none';
                }
            }
        }

        ringo_open = 1;
        ringo_save_data = data;

        setTimeout(function() { ringo_open = 2; }, 100);
    } else {
        document.getElementById(element[0]).style.display = 'none';

        ringo_open = 0
    }
}

document.addEventListener("click", function() {
    let cel_list = [];
    for(let for_a in ringo_menu_list) {
        cel_list.push(document.getElementById(ringo_menu_list[for_a]));
    }

    if(ringo_save_data !== '' && ringo_open === 2) {
        document.getElementById(ringo_save_data).style.display = 'none';

        setTimeout(function() { ringo_open = 0; }, 100);
    }
});

document.getElementById("sort_by_lang").addEventListener("click", function () {
    event.preventDefault(); // 기본 동작(새로고침) 막기
    const ul = document.querySelector("ul:nth-of-type(2)"); // 두 번째 <ul> 선택
    const items = Array.from(ul.querySelectorAll("li")); // <li> 요소 가져오기

    // 가나다 순으로 정렬
    items.sort((a, b) => {
        const aTitle = a.querySelector("a").textContent.trim();
        const bTitle = b.querySelector("a").textContent.trim();
        return aTitle.localeCompare(bTitle, "ko"); // 한국어 기준으로 정렬
    });

    // 정렬된 <li> 항목을 다시 추가
    ul.innerHTML = ""; // 기존 내용 제거
    items.forEach(item => ul.appendChild(item)); // 정렬된 항목 추가
});



document.getElementById("sort_by_views").addEventListener("click", function () {
    event.preventDefault(); // 기본 동작(새로고침) 막기
    const ul = document.querySelector("ul:nth-of-type(2)"); // 두 번째 <ul> 선택
    const items = Array.from(ul.querySelectorAll("li")); // <li> 요소 가져오기

    // 조회수를 기준으로 내림차순 정렬
    items.sort((a, b) => {
        const aViewCount = parseInt(a.querySelector("span").textContent.replace("조회수:", "").trim());
        const bViewCount = parseInt(b.querySelector("span").textContent.replace("조회수:", "").trim());
        return bViewCount - aViewCount; // 내림차순
    });

    // 정렬된 <li> 항목을 다시 추가
    ul.innerHTML = ""; // 기존 내용 제거
    items.forEach(item => ul.appendChild(item)); // 정렬된 항목 추가
});