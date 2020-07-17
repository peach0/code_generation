<?php
/**************************************************************************
 *
 * Copyright (c) 2017 Zuoyebang.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @filename:      {{.FileName}}.php
 * @desc:          {{.TableComent}}
 * @create:        {{.CreateTime}}
 * @last modified: {{.CreateTime}}
 */
class Dao_Tag extends Dao_BaseDao
{
    public function __construct() {
        parent::__construct();
        $this->_table       = '{{.TableName}}';
        $this->arrFieldsMap = [
{{.ArrFieldsMap}}
        ];
        $this->arrTypesMap  = [
{{.ArrTypesMap}}
        ];
    }
}
