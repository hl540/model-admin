console.log("model_table.js");
// 加载缓存
let columnConfig = store(tablePageKey + "_columns")
columnConfig = columnConfig == undefined ? {} : JSON.parse(columnConfig)

window.onload = function () {
    let columnSetting = columns.map((value) => {
        let newValue = {...value}
        newValue.click = tHeadSettingEvent;
        if (columnConfig[value.name] !== undefined && !columnConfig[value.name]) {
            newValue.checked = false
        }
        return newValue;
    })
    // 表头设置
    Eadmin.dropdown('#t-head-setting', {
        width: 200,
        title: '表头设置',
        data: columnSetting
    });
}

// 表头设置事件
function tHeadSettingEvent(event) {
    let col = columns[event._index];
    columnConfig[col.field] = event._checked;
    store(tablePageKey + "_columns", JSON.stringify(columnConfig));
}