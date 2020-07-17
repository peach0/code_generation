<?php
/**************************************************************************
 * Copyright (c) 2020 Zuoyebang.com, Inc. All Rights Reserved
 **************************************************************************/

/**
 * @filename:      {{.FileName}}.php
 * @desc:          {{.TableComent}}
 * @create:        {{.CreateTime}}
 * @last modified: {{.CreateTime}}
 */
class Dao_{{.FileName}} extends Wxsk_Base_Dao
{
    public static $allFields = [];

    public function __construct() {
        $this->_dbName      = 'wechat/wechat_wxsk';
        $this->_db          = null;
        $this->_table       = '{{.TableName}}';
        $this->_tableName   = '{{.TableName}}';
        $this->arrFieldsMap = [
{{.ArrFieldsMap}}
        ];
        $this->arrTypesMap  = [
{{.ArrTypesMap}}
        ];
        self::$allFields    = array_keys($this->arrFieldsMap);
    }
}
