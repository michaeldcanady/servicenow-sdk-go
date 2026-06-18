import os

files_to_fix = [
    './appserviceapi/appservice_request_builders.go',
    './attachmentapi/attachment_file_request_builder.go',
    './attachmentapi/attachment_item_file_request_builder.go',
    './attachmentapi/attachment_item_file_request_builder_test.go',
    './attachmentapi/attachment_item_request_builder.go',
    './attachmentapi/attachment_item_request_builder_test.go',
    './attachmentapi/attachment_request_builder.go',
    './attachmentapi/attachment_upload_request_builder.go',
    './batchapi/batch_request_builder_test.go',
    './cmdbinstanceapi/cmdb_class_request_builder.go',
    './cmdbinstanceapi/cmdb_item_request_builder.go',
    './cmdbinstanceapi/cmdb_relation_item_request_builder.go',
    './cmdbinstanceapi/cmdb_relation_request_builder.go',
    './documentsapi/action_request_builder.go',
    './documentsapi/attach_request_builder.go',
    './documentsapi/content_request_builder.go',
    './documentsapi/create_document_request_builder.go',
    './documentsapi/create_request_builder.go',
    './documentsapi/delete_request_builder.go',
    './documentsapi/explore_request_builder.go',
    './documentsapi/sync_down_request_builder.go',
    './documentsapi/versions_request_builder.go',
    './documentsapi/version_state_request_builder.go',
    './internal/page_iterator.go',
    './policyapi/policies_mappings_request_builder.go',
    './tableapi/table_item_request_builder.go',
    './tableapi/table_request_builder.go'
]

for file_path in files_to_fix:
    if os.path.exists(file_path):
        with open(file_path, 'r') as f:
            content = f.read()
        
        # This is a naive replacement, but should work for the standard Kiota structure
        # I'll use a simpler search/replace pattern.
        import re
        # Pattern to find the errorMapping declaration and replace it
        pattern = r'errorMapping := abstractions\.ErrorMappings\{.*?\}\s*'
        new_content = re.sub(pattern, 'errorMapping := internal.DefaultErrorMapping()\n', content, flags=re.DOTALL)
        
        if content != new_content:
            with open(file_path, 'w') as f:
                f.write(new_content)
            print(f"Fixed {file_path}")
